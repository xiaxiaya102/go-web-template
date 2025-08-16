package model

import (
	"time"
)

// BaseModel 基础模型，包含公共字段
type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

// SystemConf 系统配置结构
type SystemConf struct {
	ServerInfo ServerInfo `mapstructure:"server" json:"server" yaml:"server"`
	Log        Log        `mapstructure:"log" json:"log" yaml:"log"`
	Database   Database   `mapstructure:"database" json:"database" yaml:"database"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

// ServerInfo 服务器配置
type ServerInfo struct {
	Port     string        `mapstructure:"port" json:"port" yaml:"port"`
	Mode     string        `mapstructure:"mode" json:"mode" yaml:"mode"`
	Timeout  time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	Password string        `mapstructure:"password" json:"password" yaml:"password"` // 临时保留，后续移除
}

// Log 日志配置
type Log struct {
	Path  string `mapstructure:"path" json:"path" yaml:"path"`
	Level string `mapstructure:"level" json:"level" yaml:"level"`
}

// Database 数据库配置
type Database struct {
	Type     string `mapstructure:"type" json:"type" yaml:"type"` // mysql, postgres, sqlite
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"` // sqlite文件路径
	SSLMode  string `mapstructure:"ssl_mode" json:"ssl_mode" yaml:"ssl_mode"`
}

// Redis 配置
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database int    `mapstructure:"database" json:"database" yaml:"database"`
}

// JWT 配置
type JWT struct {
	Secret     string        `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expire     time.Duration `mapstructure:"expire" json:"expire" yaml:"expire"`
	RefreshTTL time.Duration `mapstructure:"refresh_ttl" json:"refresh_ttl" yaml:"refresh_ttl"`
}
