package resource

import (
	"github.com/Haski007/fav-videos/internal/fvb/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type FVBService struct {
	Bot *tgbotapi.BotAPI
	Cfg *config.Bot
}

func NewFVBService(cfg *config.Bot) (*FVBService, error) {
	var err error

	bot := &FVBService{}
	bot.Cfg = cfg
	bot.Bot, err = tgbotapi.NewBotAPI(bot.Cfg.GetToken().String())
	if err != nil {
		return nil, err
	}

	bot.Bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot, nil
}
