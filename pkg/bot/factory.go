package bot

func NewBot(botType string, webhook string) Bot {
	switch botType {
	//case "wecom":
	//	return &WeComBot{WebhookURL: webhook}
	//case "dingtalk":
	//	return &DingTalkBot{WebhookURL: webhook}
	case "feishu":
		return &FeiShuBot{WebhookURL: webhook}
	default:
		return nil
	}
}
