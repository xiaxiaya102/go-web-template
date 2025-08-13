package bot

// Bot 是所有机器人的统一接口
type Bot interface {
	SendMessage(content string) error
}
