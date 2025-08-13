package service

//
//import (
//	"encoding/json"
//	"feishuReboot/global"
//	"feishuReboot/logger"
//	"fmt"
//	"github.com/go-resty/resty/v2"
//)
//
//// 定义平台代码和平台名称为常量
//const (
//	Baidu         = "baidu"
//	Shaoshupai    = "shaoshupai"
//	Weibo         = "weibo"
//	Zhihu         = "zhihu"
//	Pojie         = "52pojie"
//	Bilibili      = "bilibili"
//	Douban        = "douban"
//	Hupu          = "hupu"
//	Tieba         = "tieba"
//	Juejin        = "juejin"
//	Douyin        = "douyin"
//	V2ex          = "v2ex"
//	JRTT          = "jinritoutiao"
//	StackOverflow = "stackoverflow"
//	GitHub        = "github"
//	HackerNews    = "hackernews"
//)
//
//// News 定义新闻结构体
//type News struct {
//	Title       string `json:"title"`
//	URL         string `json:"url"`
//	Content     string `json:"content"`
//	Source      string `json:"source"`
//	PublishTime string `json:"publish_time"`
//}
//
//// Response 定义响应结构体
//type Response struct {
//	Status string `json:"status"`
//	Data   []News `json:"data"`
//}
//
//// GetLastLyNewsByPlatformCode 根据平台编码获取平台热点新闻
//func GetLastLyNewsByPlatformCode(code string) string {
//	client := resty.New()
//
//	resp, err := client.R().Get(global.System.NewsUrl + code)
//	if err != nil {
//		logger.Error("请求失败")
//	}
//	//使用 map 来解析 JSON 数据
//	var response Response
//	err = json.Unmarshal(resp.Body(), &response)
//	if err != nil {
//		logger.Error("Error unmarshaling JSON:", err)
//	}
//
//	status := response.Status
//	if status == "200" {
//		news := response.Data
//		logger.Info("热点新闻获取成功,供获取到:%d条", len(news))
//		return formatMultipleNewsMessages(news)
//	}
//	return ""
//}
//
//// formatNewsMessage 格式化单条新闻消息
//func formatNewsMessage(index int, news News) string {
//	return fmt.Sprintf(
//		"序号: %d\n"+
//			"标题: %s\n"+
//			"来源: %s\n"+
//			"发布时间: %s\n"+
//			"详情: (%s)\n",
//		index, news.Title, news.Source, news.PublishTime, news.URL,
//	)
//}
//
//// formatMultipleNewsMessages 格式化多条新闻消息
//func formatMultipleNewsMessages(newsList []News) string {
//	var message string
//	for i, news := range newsList {
//		message += formatNewsMessage(i, news) + "\n"
//	}
//	return message
//}
//
//// sendFeishuMessage 发送消息到飞书群组
//func sendFeishuMessage(message string) {
//	client := resty.New()
//
//	// 构造发送的 JSON 数据
//	payload := map[string]interface{}{
//		"msg_type": "text",
//		"content": map[string]string{
//			"text": "每日热点新闻\n\n" + message,
//		},
//	}
//
//	result, err := client.R().SetBody(payload).Post(global.System.FeishuWebhook)
//	if err != nil {
//		return
//	}
//	logger.Info("消息发送结果:%s", result)
//}
