package config

import (
	"feishuReboot/global"
	"feishuReboot/logger"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configurePath = "config" // 配置文件所在目录
)

var (
	Conf           Config
	lastReloadTime time.Time
	reloadMutex    sync.Mutex
)

func LoadConfig() {
	Conf = Config{}
	Conf.v = viper.New()

	v := Conf.Application.v
	v.SetConfigFile("config/application-web.yaml")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("load config path: %s, error: %s", configurePath, err)
	}

	if err := v.Unmarshal(&global.System); err != nil {
		fmt.Printf("Unable to decode into struct, %s", err)
	}

	// 启用配置文件热更新
	watchConfig(v, func(e fsnotify.Event) {
		// 防止重复触发
		reloadMutex.Lock()
		defer reloadMutex.Unlock()

		now := time.Now()
		if now.Sub(lastReloadTime) < 100*time.Millisecond {
			return // 100ms内的重复事件忽略
		}
		lastReloadTime = now

		logger.Info("配置文件发生变化: %s", e.Name)

		// 加锁防止并发读写
		Conf.Lock()
		defer Conf.Unlock()

		// 重新解析配置到全局变量
		if err := v.Unmarshal(&global.System); err != nil {
			logger.Error("热更新配置失败: %v", err)
		} else {
			logger.Info("配置热更新成功")
		}
	})
}

type Application struct {
	name string
	v    *viper.Viper
}

func (c *Application) GetName() string {
	return c.name
}

func (c *Application) GetViper() *viper.Viper {
	return c.v
}

type Config struct {
	Application
	rwMutex sync.RWMutex
}

func (sc *Config) RLock() {
	sc.rwMutex.RLock()
}

func (sc *Config) RUnlock() {
	sc.rwMutex.RUnlock()
}

func (sc *Config) Lock() {
	sc.rwMutex.Lock()
}

func (sc *Config) Unlock() {
	sc.rwMutex.Unlock()
}

// 监控配置文件变化并热加载程序
func watchConfig(v *viper.Viper, callback func(in fsnotify.Event)) {
	v.OnConfigChange(callback)
	v.WatchConfig()
}

func SaveConfig() error {
	out, err := yaml.Marshal(global.System)
	if err != nil {
		logger.Error("error marshalling yaml: %v", err)
		return err
	}

	err = os.WriteFile("config/application-web.yaml", out, 0644) // 0644 是文件权限
	if err != nil {
		logger.Error("error marshalling yaml: %v", err)
		return err
	}

	return nil
}
