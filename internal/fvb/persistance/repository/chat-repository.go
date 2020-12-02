package repository

import "errors"

var (
	ErrChatAlreadyExists = errors.New("CHAT_ALREADY_EXISTS")
	ErrChatDoesNotExist  = errors.New("CHAT_DOES_NOT_EXIST")
)

type ChatRepository interface {
	SaveNewChat()
}
