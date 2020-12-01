package resource

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (bot *FVBService) Reply(update tgbotapi.Update, message string) {
	resp := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	_, err := bot.Bot.Send(resp)
	if err != nil {
		logrus.Printf("[send message /help] err: %s", err)
		return
	}
}
