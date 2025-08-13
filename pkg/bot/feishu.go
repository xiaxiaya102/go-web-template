package bot

import (
	"feishuReboot/logger"
	"github.com/go-resty/resty/v2"
)

type FeiShuBot struct {
	WebhookURL string
}

func (b *FeiShuBot) SendMessage(content string) error {
	client := resty.New()

	// 构造发送的 JSON 数据
	payload := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": "每日热点新闻\n\n" + content,
		},
	}

	result, err := client.R().SetBody(payload).Post(b.WebhookURL)
	if err != nil {
		return nil
	}
	logger.Debug("消息发送结果:%s", result)
	return nil
}
