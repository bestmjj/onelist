package tg

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bestmjj/onelist/onelist/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// PostTelegramMessage 函数用于向指定的聊天 ID 发送消息
func PostTelegramMessage(message string) error {
	// 初始化 Telegram Bot
	bot, err := tgbotapi.NewBotAPI(config.TGBotKey)
	if err != nil {
		return fmt.Errorf("创建 Telegram Bot 失败: %v", err)
	}

	// 设置调试模式（可选）
	bot.Debug = true
	log.Printf("已授权账户 %s", bot.Self.UserName)

	// 创建一个新的消息
	id, err := convertStringToInt64(config.TGChatID)
	if err != nil {
		return fmt.Errorf("转换 Chat ID 失败: %v", err)
	}
	msg := tgbotapi.NewMessage(id, message)
	msg.ParseMode = "HTML" // 使用 HTML 格式化消息内容

	// 发送消息
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}

	return nil
}

func convertStringToInt64(str string) (int64, error) {
	// 将字符串转换为 int64
	return strconv.ParseInt(str, 10, 64)
}
