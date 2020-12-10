package resource

import (
	"fmt"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/model"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository"
	"github.com/Haski007/fav-videos/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandPublishersHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var publishers []*model.Publisher

	if err := bot.ChatRepository.GetAllPublishers(chatID, &publishers); err != nil {
		switch err {
		case repository.ErrChatDoesNotExist:
			bot.Reply(
				chatID,
				"This chat is not registered yet! "+emoji.NoEntry)
			return
		default:
			bot.Reply(
				chatID,
				"Internal Error! "+emoji.NoEntry)
			bot.ReportToTheCreator(
				fmt.Sprintf(
					"Error [commandDelPublisherHandler] mongo.RemovePublisher chatID: %+v \n err: %s", chatID, err),
			)
			return
		}
	}

	var answer string

	for i, p := range publishers {
		answer += fmt.Sprintf("%-2d) %s\n", i+1, p.Username)
	}

	bot.Reply(chatID, answer)
}
