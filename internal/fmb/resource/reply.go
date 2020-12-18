package resource

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (bot *FVBService) Reply(chatID int64, message string) {
	resp := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Bot.Send(resp)
	if err != nil {
		logrus.Printf("[send message /help] err: %s", err)
		return
	}
}
