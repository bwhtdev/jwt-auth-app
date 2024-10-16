package storage

import (
	"github.com/google/uuid"
	"database/sql"
	types "backend/pkg/types"
)

func (s *DBStore) CreatePerson(person *types.People) (uuid.UUID, error) {
	var id uuid.UUID
	
	query := `INSERT INTO people
		(name)
		VALUES ($1)
		RETURNING id;`
	err := s.db.QueryRow(query, person.Name).Scan(&id)
	if err != nil {
		return uuid.New(), err
	}
	return id, nil
}

func (s *DBStore) GetPeople() ([]*types.People, error) {
	rows, err := s.db.Query("SELECT * FROM people;")
	if err != nil {
		return nil, err
	}

	people := []*types.People{}
	for rows.Next() {
		person, err := scanIntoPerson(rows)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func scanIntoPerson(rows *sql.Rows) (*types.People, error) {
	person := new(types.People)
	err := rows.Scan(
		&person.ID,
		&person.Name,
		&person.CreatedAt)
	return person, err
}
