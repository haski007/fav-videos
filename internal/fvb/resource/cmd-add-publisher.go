package resource

import (
	"fmt"
	"strings"

	"github.com/Haski007/fav-videos/api"

	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository"

	"github.com/Haski007/fav-videos/pkg/emoji"

	"github.com/Haski007/fav-videos/internal/fvb/persistance/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandAddPublisherHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) == 0 {
		bot.Reply(
			chatID,
			"No username.\nPlease use command like that:\n/add_publisher [tiktok-username]")
		return
	}

	username := args[0]
	secID, err := api.GetSecureUserID(username)
	if err != nil || secID == "" {
		bot.Reply(
			chatID,
			"Your username is invalid! "+emoji.NoEntry)
		return
	}
	publisher := model.NewPublisher(args[0], secID)

	if err := bot.ChatRepository.PushNewPublusher(chatID, publisher); err != nil {
		switch err {
		case repository.ErrChatDoesNotExist:
			bot.Reply(
				chatID,
				"Your chat is not registered!")
			return
		default:
			bot.Reply(
				chatID,
				"Internal Error! "+emoji.NoEntry)
			bot.ReportToTheCreator(fmt.Sprintf("Error [PushNewPublisher] pub: %+v \n err: %s", publisher, err))
			return
		}

	}

	bot.Reply(chatID, "Publisher added! "+emoji.Check)
}
