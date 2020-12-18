package mongodb

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"

	"github.com/Haski007/fav-videos/api"

	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/model"
	"github.com/Haski007/fav-music-bot/internal/fvb/persistance/repository"

	"github.com/stretchr/testify/assert"
)

const (
	mongoAddr = "mongodb://root:demian196@172.20.0.2:27017"
)

func getCollection() *mgo.Collection {
	session, err := mgo.Dial(mongoAddr)
	if err != nil {
		logrus.Fatalf("[mgo Dial] addr: %s | err: %s", cfg.Addr, err)
	}

	if err = session.Ping(); err != nil {
		logrus.Fatalf("[mgo Ping] addr: %s | err: %s", cfg.Addr, err)
	}

	return session.DB(cfg.DBName).C(collName)
}

func TestChatRepository_PushNewPublusher(t *testing.T) {
	t.Skip()
	rep := &ChatRepository{
		coll: getCollection(),
	}

	username := "demianchik"
	secID, _ := api.GetSecureUserID(username)
	publisher := model.NewPublisher(username, secID)

	tt := []struct {
		name      string
		chatID    int64
		publisher *model.Publisher
		expectErr error
	}{
		{
			name:      "Test with username: demianchik",
			chatID:    370649141,
			publisher: publisher,
			expectErr: nil,
		},
		{
			name:      "Test error",
			chatID:    3706491412,
			publisher: publisher,
			expectErr: repository.ErrChatDoesNotExist,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := rep.PushNewPublusher(tc.chatID, tc.publisher)
			assert.Equal(t, tc.expectErr, err)
		})
	}
}
