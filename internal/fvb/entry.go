package fvb

import (
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/Haski007/fav-videos/pkg/factory"
	"github.com/Haski007/fav-videos/pkg/run"
)

func Run(args *run.Args) error {
	var conf config.Config

	factory.InitLog(args.LogLevel)

	app, err := NewService(args, &conf)
	return nil
}
