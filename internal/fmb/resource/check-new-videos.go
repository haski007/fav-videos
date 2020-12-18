package resource

import (
	"fmt"
	"log"
	"time"

	"github.com/Haski007/fav-music-bot/pkg/emoji"
	"github.com/Haski007/fav-music-bot/api"
	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/model"
	"github.com/Haski007/fav-music-bot/pkg/file"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FVBService) CheckNewVideos() {
	defer func() {
		if recoveryErr := recover(); recoveryErr != nil {
			message := fmt.Sprintf("Panic [CheckNewVideos] err: %g", recoveryErr)
			bot.ReportToTheCreator(message)
			logrus.Errorf(message)
		}
	}()

	var chats []model.Chat
	bot.ChatRepository.GetAllChats(&chats)

	for _, chat := range chats {
		for _, publisher := range chat.Publishers {
			likes, err := api.GetLikedVideos(publisher, 3)
			if err != nil {
				logrus.Printf("[GetLikedVideos] err: %s", err)
				return
			}

			var tmpFileName = ".tmp_video.mp4"

			for _, v := range likes {
				if bot.ChatRepository.PostedVideoExists(chat.ID, v.ID) {
					continue
				}

				log.Println("Posting", v.ID)

				if err := file.GetFileByUrl(tmpFileName, v.DownloadURL); err != nil {
					logrus.Errorf("[file GetFileByUrl] err: %s", err)
				}

				resp := tgbotapi.NewVideoUpload(chat.ID, tmpFileName)
				resp.Caption = fmt.Sprintf("Liked by %s %s!", v.Publisher, emoji.Heart)
				resp.ReplyMarkup = model.NewOriginalURLMarkup(v.ShareURL)

				if message, err := bot.Bot.Send(resp); err != nil {
					logrus.Errorf("[bot Send] video | message: %+v | err: %s", message, err)
				}
				if err := bot.ChatRepository.PushPostedVideo(chat.ID, v.ID); err != nil {
					bot.Reply(chat.ID, "Internal Error! "+emoji.NoEntry)
					bot.ReportToTheCreator(
						fmt.Sprintf("[PushPostedVideo] chatID: %d | videoID: %s", chat.ID, v.ID))
				}
				time.Sleep(time.Second * 3)
			}
		}
	}
}
