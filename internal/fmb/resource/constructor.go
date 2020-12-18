package resource

import (
	"github.com/Haski007/fav-music-bot/internal/fvb/config"
	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/repository/mongodb"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type FVBService struct {
	Bot            *tgbotapi.BotAPI
	Cfg            *config.Bot
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
