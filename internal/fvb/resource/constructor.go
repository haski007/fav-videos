package resource

import (
	"errors"
	"fmt"
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/caarlos0/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

type FVBService struct {
	Bot       *tgbotapi.BotAPI
	Cfg       *config.Bot
	TiktokCfg *config.TikTok
}

func NewFVBService() (*FVBService, error) {
	var err error

	bot := &FVBService{}

	// ---> Bot configs
	bot.Cfg = &config.Bot{}
	if err := env.Parse(bot.Cfg); err != nil {
		logrus.Fatalf("[env Parse] Bot config err: %s", err)
	}

	// ---> TikTok configs
	bot.TiktokCfg = &config.TikTok{}
	if err := env.Parse(bot.TiktokCfg); err != nil {
		logrus.Fatalf("[env Parse] TikTok config err: %s", err)
	}
	if bot.TiktokCfg.SecUserID == "" {
		bot.TiktokCfg.SecUserID, err = getSecUserID(bot.TiktokCfg.Username)
		if err != nil {
			log.Fatalln("SecUID", err)
		}
	}
	fmt.Println(bot.TiktokCfg.SecUserID)

	bot.Bot, err = tgbotapi.NewBotAPI(bot.Cfg.GetToken().String())
	if err != nil {
		return nil, err
	}

	bot.Bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot, nil
}

func getSecUserID(username string) (string, error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}

	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI("https://www.tiktok.com/@" + username)
	req.Header.SetUserAgent(config.UserAgent)

	err := fasthttp.Do(req, res)
	if err != nil {
		return "", err
	}

	matches := config.SecUIDReg.FindStringSubmatch(res.String())
	if len(matches) != 2 {
		return "", errors.New("no matches")
	}

	return matches[1], nil
}
