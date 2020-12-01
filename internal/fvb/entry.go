package fvb

import (
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/Haski007/fav-videos/internal/fvb/resource"
	"github.com/Haski007/fav-videos/pkg/factory"
	"github.com/Haski007/fav-videos/pkg/run"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

func Run(args *run.Args) error {
	cfg := config.Bot{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("[env Parse] err: %s", err)
	}

	botService, err := resource.NewFVBService(&cfg)
	if err != nil {
		logrus.Fatalf("[NewFVBService] err: %s", err)
	}

	factory.InitLog(args.LogLevel)

	StartBot(botService)
	return nil
}
