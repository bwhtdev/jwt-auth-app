package types

import (
	"time"
	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateMessageRequest struct {
	Text      string    `json:"text"`
	Username  string    `json:"username"`
}

type UpdateMessageRequest struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
}

type DeleteMessageRequest struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
}

func NewMessage(text, username string) (*Message, error) {
	return &Message{
		Text: text,
		Username: username,
		CreatedAt: time.Now().UTC(),
	}, nil
}
