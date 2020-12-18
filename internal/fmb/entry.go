package fvb

import (
	"github.com/Haski007/fav-music-bot/internal/fvb/resource"
	"github.com/Haski007/fav-music-bot/pkg/factory"
	"github.com/Haski007/fav-music-bot/pkg/run"
	"github.com/sirupsen/logrus"
)

func Run(args *run.Args) error {
	botService, err := resource.NewFVBService()
	if err != nil {
		logrus.Fatalf("[NewFVBService] err: %s", err)
	}

	factory.InitLog(args.LogLevel)

	StartBot(botService)
	return nil
}
