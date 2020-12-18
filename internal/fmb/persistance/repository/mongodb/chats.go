package mongodb

import (
	"github.com/Haski007/fav-music-bot/internal/fvb/config"
	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/model"
	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/repository"
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

// ---> CHATS

func (r *ChatRepository) GetAllChats(chats *[]model.Chat) {
	if err := r.coll.Find(bson.M{}).All(chats); err != nil {
		logrus.Errorf("[GetAllChats] err: %s", err)
		return
	}
}

func (r *ChatRepository) SaveNewChat(chat *model.Chat) error {
	if r.ChatExists(chat.ID) {
		return repository.ErrChatAlreadyExists
	}

	return r.coll.Insert(chat)
}

func (r *ChatRepository) RemoveChat(chatID int64) error {

	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	return r.coll.RemoveId(chatID)
}

func (r *ChatRepository) ChatExists(id int64) bool {
	count, _ := r.coll.FindId(id).Count()
	if count != 0 {
		return true
	}
	return false
}

// ---> Users

func (r *ChatRepository) UserExists(chatID int64, username string) bool {
	query := bson.M{
		"_id":                 chatID,
		"publishers.username": username,
	}

	count, _ := r.coll.Find(query).Count()
	if count != 0 {
		return true
	}
	return false
}

func (r *ChatRepository) PushNewPublusher(chatID int64, pub *model.Publisher) error {
	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	if r.UserExists(chatID, pub.Username) {
		return repository.ErrUserAlreadyExists
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

func (r *ChatRepository) RemovePublisher(chatId int64, username string) error {

	if !r.UserExists(chatId, username) {
		return repository.ErrUserDoesNotExist
	}

	findQuery := bson.M{
		"_id": chatId,
	}
	updateQuery := bson.M{
		"$pull": bson.M{
			"publishers": bson.M{
				"username": username,
			},
		},
	}

	return r.coll.Update(findQuery, updateQuery)
}

func (r *ChatRepository) GetAllPublishers(chatID int64, publishers *[]*model.Publisher) error {
	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	var chat model.Chat

	if err := r.coll.FindId(chatID).One(&chat); err != nil {
		return err
	}

	pubs := make([]*model.Publisher, len(chat.Publishers))
	for i, p := range chat.Publishers {
		pubs[i] = p
	}

	*publishers = pubs
	return nil
}

// ---> Videos

func (r *ChatRepository) PushPostedVideo(chatID int64, videoID string) error {
	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	findQuery := bson.M{
		"_id": chatID,
	}
	updateQuery := bson.M{
		"$push": bson.M{
			"posted_videos": videoID,
		},
	}

	err := r.coll.Update(findQuery, updateQuery)
	return err
}

func (r *ChatRepository) PostedVideoExists(chatID int64, videoID string) bool {
	query := bson.M{
		"_id":           chatID,
		"posted_videos": videoID,
	}
	count, _ := r.coll.Find(query).Count()
	if count != 0 {
		return true
	}
	return false
}
