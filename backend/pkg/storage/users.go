package storage

import (
	"database/sql"
	"fmt"

	types "backend/pkg/types"
)

func (s *DBStore) CreateUser(user *types.User) error {
	query := `INSERT INTO users
	(username, encrypted_password)
	VALUES ($1, $2);`

	_, err := s.db.Query(
		query,
		user.Username,
		user.EncryptedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (s *DBStore) UpdateUser(user *types.User) error {
	query := `UPDATE users
	SET username=$2
	WHERE id=$1;`

	_, err := s.db.Query(
		query,
		user.ID,
		user.Username)

	if err != nil {
		return err
	}

	return nil
}

func (s *DBStore) DeleteUser(id string) error {
	_, err := s.db.Query("DELETE FROM users WHERE id = $1", id)
	return err
}

func scanIntoUser(rows *sql.Rows) (*types.User, error) {
	user:= new(types.User)
	err := rows.Scan(
	    &user.ID,
		&user.Username,
		&user.EncryptedPassword,
	    &user.CreatedAt)

	return user, err
}

/*func (s *PostgresStore) GetUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("users not found")
}*/

func (s *DBStore) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user with username [%s] not found", username)
}

func (s *DBStore) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}
