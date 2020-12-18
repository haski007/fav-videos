package resource

import (
	"fmt"
	"strings"

	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/repository"
	"github.com/Haski007/fav-music-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandDelPublisherHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) == 0 {
		bot.Reply(
			chatID,
			"No username.\nPlease use command like that:\n/del_publisher [tiktok-username]")
		return
	}

	username := args[0]

	if err := bot.ChatRepository.RemovePublisher(chatID, username); err != nil {
		switch err {
		case repository.ErrUserDoesNotExist:
			bot.Reply(
				chatID,
				"This user does not exist! "+emoji.NoEntry)
			return
		default:
			bot.Reply(
				chatID,
				"Internal Error! "+emoji.NoEntry)
			bot.ReportToTheCreator(
				fmt.Sprintf(
					"Error [commandDelPublisherHandler] mongo.RemovePublisher pub: %+v \n err: %s", username, err),
			)
			return
		}

	}

	bot.Reply(chatID, "Publisher removed! "+emoji.Basket)
}
