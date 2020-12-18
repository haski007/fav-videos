package resource

import (
	"fmt"

	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository"
	"github.com/Haski007/fav-videos/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) commandDelChatHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := bot.ChatRepository.RemoveChat(chatID); err != nil {
		switch err {
		case repository.ErrChatDoesNotExist:
			bot.Reply(chatID, "Chat is not registered! "+emoji.Failed)
		default:
			bot.Reply(chatID, "Chat is not registered! "+emoji.Failed)
			bot.ReportToTheCreator(fmt.Sprintf("[RemoveChat] chatID: %d | err: %s", chatID, err))
			return
		}
	}

	bot.Reply(chatID, "Chat has been removed! "+emoji.Basket)
}
