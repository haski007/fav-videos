package resource

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandHelpHandler(update tgbotapi.Update) {
	bot.Reply(
		update.Message.Chat.ID,
		"Here is bot to subscribe on someone's liked videos in TikTok\n"+
			"You can use such commands:\n"+
			" - /help - for help\n"+
			" - /reg_chat - for adding current chat to DataBase\n"+
			" - /del_chat - for removing chat\n"+
			" - /add_publisher [tiktok username] - to add new publisher\n"+
			" - /del_publisher [tiktok username] - to remove publisher\n"+
			" - /publishers - to get list of all publishers\n")
}
