package resource

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandHelpHandler(update tgbotapi.Update) {
	bot.Reply(
		update,
		"Here is bot to subscribe on someone's liked videos in TikTok")
}
