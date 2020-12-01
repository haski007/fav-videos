package resource

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (bot *FVBService) HandleRoutes(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.CommandWithAt() {
			case "help":
				go bot.commandHelpHandler(update)
			}
		}
	}
}
