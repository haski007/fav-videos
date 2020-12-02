package resource

import (
	"github.com/Haski007/fav-videos/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) HandleRoutes(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.CommandWithAt() {
			case "help":
				go bot.commandHelpHandler(update)
			case "reg_chat":
				go bot.regNewChat(update)
			case "add_publisher":
				go bot.addPublisher(update)
			default:
				bot.Reply(update.Message.Chat.ID, "Such command does not exist! "+emoji.NoEntry)
			}
		}
	}
}
