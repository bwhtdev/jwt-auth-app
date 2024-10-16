package types

import (
	"time"
	"github.com/google/uuid"
)

type People struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreatePeopleRequest struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPerson(name string) (*People, error) {
	return &People{
		Name: name,
		CreatedAt: time.Now().UTC(),
	}, nil
}
