package resource

import (
	"fmt"

	"github.com/Haski007/fav-videos/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) HandleRoutes(updates tgbotapi.UpdatesChannel) {
	botCreds, err := bot.Bot.GetMe()
	if err != nil {
		bot.ReportToTheCreator(
			fmt.Sprintf("[bot GetMe] err: %s", err))
		return
	}
	for update := range updates {
		// Check if someone removes bot
		if update.Message.LeftChatMember != nil &&
			update.Message.LeftChatMember.UserName == botCreds.UserName {
			bot.ChatRepository.RemoveChat(update.Message.Chat.ID)
		}

		if command := update.Message.CommandWithAt(); command != "" {
			switch {
			case command == "help" || command == "help"+"@"+botCreds.UserName:
				go bot.commandHelpHandler(update)
			case command == "reg_chat" || command == "reg_chat"+"@"+botCreds.UserName:
				go bot.commandRegNewChatHandler(update)
			case command == "del_chat" || command == "del_chat"+"@"+botCreds.UserName:
				go bot.commandDelChatHandler(update)
			case command == "add_publisher" || command == "add_publisher"+"@"+botCreds.UserName:
				go bot.commandAddPublisherHandler(update)
			case command == "del_publisher" || command == "del_publisher"+"@"+botCreds.UserName:
				go bot.commandDelPublisherHandler(update)
			case command == "publishers" || command == "publishers"+"@"+botCreds.UserName:
				go bot.commandPublishersHandler(update)
			default:
				bot.Reply(update.Message.Chat.ID, "Such command does not exist! "+emoji.NoEntry)
			}
		}
	}
}
