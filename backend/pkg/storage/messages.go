package storage

import (
	"github.com/google/uuid"
	"database/sql"
	types "backend/pkg/types"
)

func (s *DBStore) CreateMessage(message *types.Message) (uuid.UUID, error) {
	var id uuid.UUID
	
	query := `INSERT INTO messages
		(text, username)
		VALUES ($1, $2)
		RETURNING id;`
	err := s.db.QueryRow(query, message.Text, message.Username).Scan(&id)
	if err != nil {
		return uuid.New(), err
	}
	return id, nil
}

func (s *DBStore) GetMessages() ([]*types.Message, error) {
	rows, err := s.db.Query("SELECT * FROM messages;")
	if err != nil {
		return nil, err
	}

	messages := []*types.Message{}
	for rows.Next() {
		message, err := scanIntoMessage(rows)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func scanIntoMessage(rows *sql.Rows) (*types.Message, error) {
	message := new(types.Message)
	err := rows.Scan(
		&message.ID,
		&message.Text,
		&message.Username,
		&message.CreatedAt)
	return message, err
}
