package resource

import (
	"log"

	"github.com/Haski007/fav-videos/api"

	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository/mongodb"
	"github.com/caarlos0/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type FVBService struct {
	Bot            *tgbotapi.BotAPI
	Cfg            *config.Bot
	TiktokCfg      *config.TikTok
	ChatRepository *mongodb.ChatRepository
}

func NewFVBService() (*FVBService, error) {
	var err error

	bot := &FVBService{}

	/*
	** ---> Bot configs
	 */
	bot.Cfg = &config.Bot{}
	if err := env.Parse(bot.Cfg); err != nil {
		logrus.Fatalf("[env Parse] Bot config err: %s", err)
	}

	/*
	** ---> TikTok configs
	 */
	bot.TiktokCfg = &config.TikTok{}
	if err := env.Parse(bot.TiktokCfg); err != nil {
		logrus.Fatalf("[env Parse] TikTok config err: %s", err)
	}
	if bot.TiktokCfg.SecUserID == "" {
		bot.TiktokCfg.SecUserID, err = api.GetSecureUserID(bot.TiktokCfg.Username)
		if err != nil {
			log.Fatalln("SecUID", err)
		}
	}

	/*
	** ---> mongo Collection
	 */
	bot.ChatRepository = &mongodb.ChatRepository{}
	bot.ChatRepository.InitChatsConn()

	bot.Bot, err = tgbotapi.NewBotAPI(bot.Cfg.GetToken().String())
	if err != nil {
		return nil, err
	}

	bot.Bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot, nil
}
