package resource

import (
	"github.com/Haski007/fav-videos/api"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/model"
	"github.com/Haski007/fav-videos/pkg/file"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func (bot *FVBService) CheckNewVideos() {
	likes, err := api.GetLikedVideos(bot.TiktokCfg.SecUserID, 10)
	if err != nil {
		logrus.Printf("[GetLikedVideos] err: %s", err)
		return
	}

	// TODO: fix crutch
	var tmpFileName = ".tmp_video.mp4"

	for _, v := range likes {
		// TODO: add collection with videos
		//if wasAlreadyPosted(v.ID) {
		//	continue
		//}

		log.Println("Posting", v.ID)

		if err := file.GetFileByUrl(tmpFileName, v.DownloadURL); err != nil {
			logrus.Errorf("[file GetFileByUrl] err: %s", err)
		}

		resp := tgbotapi.NewVideoUpload(bot.Cfg.CreatorID, tmpFileName)
		resp.ReplyMarkup = model.NewOriginalURLMarkup(v.ShareURL)

		if message, err := bot.Bot.Send(resp); err != nil {
			logrus.Errorf("[bot Send] video | message: %+v | err: %s", message, err)
		}

		time.Sleep(time.Second * 3)
	}
}
