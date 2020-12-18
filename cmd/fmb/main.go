package main

import (
	"fmt"
	"os"

	"github.com/Haski007/fav-music-bot/internal/fvb"
	"github.com/Haski007/fav-music-bot/pkg/run"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.App{
		Name:    "fav-videos-fvb",
		Usage:   "Favourite tiktok videos tg fvb entry point",
		Version: Version,
		Action: func(c *cli.Context) error {
			if err := fvb.Run(&run.Args{
				LogLevel: run.LogLevel(c.String("info")),
			}); err != nil {
				return fmt.Errorf("run: %w", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
