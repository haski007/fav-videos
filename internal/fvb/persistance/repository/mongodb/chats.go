package mongodb

import (
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/model"
	"github.com/Haski007/fav-videos/internal/fvb/persistance/repository"
	"github.com/caarlos0/env"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

const (
	collName = "chats"
)

var session *mgo.Session
var cfg config.MongoCfg

type ChatRepository struct {
	coll *mgo.Collection
}

func (r *ChatRepository) InitChatsConn() {
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("[env Parse] MongoCfg err: %s", err)
	}

	cfg.Addr = "mongodb://" + cfg.Username + ":" + cfg.Password + "@" + cfg.HostName + ":" + cfg.Port

	session, err := mgo.Dial(cfg.Addr)
	if err != nil {
		logrus.Fatalf("[mgo Dial] addr: %s | err: %s", cfg.Addr, err)
	}

	if err = session.Ping(); err != nil {
		logrus.Fatalf("[mgo Ping] addr: %s | err: %s", cfg.Addr, err)
	}

	r.coll = session.DB(cfg.DBName).C(collName)
}

func (r *ChatRepository) SaveNewChat(chat *model.Chat) error {
	if r.ChatExists(chat.ID) {
		return repository.ErrChatAlreadyExists
	}

	return r.coll.Insert(chat)
}

func (r *ChatRepository) PushNewPublusher(chatID int64, pub *model.Publisher) error {
	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	findQuery := bson.M{
		"_id": chatID,
	}
	updateQuery := bson.M{
		"$push": bson.M{
			"publishers": pub,
		},
	}

	err := r.coll.Update(findQuery, updateQuery)
	return err
}

func (r *ChatRepository) ChatExists(id int64) bool {
	count, _ := r.coll.FindId(id).Count()
	if count != 0 {
		return true
	}
	return false
}
