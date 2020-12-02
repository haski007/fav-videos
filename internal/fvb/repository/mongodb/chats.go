package mongodb

import (
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/caarlos0/env"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	collName = "chats"
)

var session *mgo.Session
var cfg config.MongoCfg

func init() {
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("[env Parse] MongoCfg err: %s", err)
	}

	session, err := mgo.Dial("mongodb://" + cfg.Username + ":" + cfg.Password + "@" + cfg.HostName + ":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}

func InitChatsConn() *mgo.Collection {
	return session.DB(cfg.DBName).C(collName)
}
