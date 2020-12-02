package resource

import (
	"fmt"

	"github.com/Haski007/fav-videos/pkg/emoji"

	"github.com/Haski007/fav-videos/internal/fvb/persistance/model"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (bot *FVBService) regNewChat(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	chatName := update.Message.Chat.FirstName + " " + update.Message.Chat.LastName

	chat := model.NewChat(chatID, chatName)

	if err := bot.ChatRepository.SaveNewChat(chat); err != nil {
		switch err {
		case repository.ErrChatAlreadyExists:
			bot.Reply(chatID, "Chat is already registered "+emoji.NoEntry)
			return
		default:
			bot.Reply(chatID, "Internal Server Error! err: %s"+err.Error())
			bot.ReportToTheCreator(fmt.Sprintf("[SaveNewChat] chat %+v | err: %s", chat, err))
			logrus.Errorf("[SaveNewChat] chat %+v | err: %s", chat, err)
			return
		}
	}

	bot.Reply(chatID, "Your chat has been registered! "+emoji.Check)
}
