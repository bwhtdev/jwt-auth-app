package storage

import (
	"fmt"
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

func (s *DBStore) UpdateMessage(message *types.UpdateMessageRequest) error {	
	query := `UPDATE messages
		SET text = $2
		WHERE id = $1;`
	_, err := s.db.Exec(query, message.ID, message.Text)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBStore) DeleteMessage(id string) error {	
	_, err := s.db.Query("DELETE FROM messages WHERE id = $1;", id)
	return err
}

func (s *DBStore) GetMessage(id string) (*types.Message, error) {
	fmt.Println(id)
	query := `SELECT * FROM messages
		WHERE id = $1;`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		return scanIntoMessage(rows)
	}
	
	return nil, fmt.Errorf("message %s not found", id)
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
