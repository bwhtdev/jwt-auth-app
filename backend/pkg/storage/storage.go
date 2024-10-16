package storage

import (
	"database/sql"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"log"

	_ "github.com/lib/pq"
	types "backend/pkg/types"
)

type Storage interface {
	CreateUser(*types.User) error
	GetUserByID(string) (*types.User, error)
	GetUserByUsername(string) (*types.User, error)
	//GetUsers() ([]*types.User, error)
	UpdateUser(*types.User) error
	DeleteUser(string) error
	
	GetPeople() ([]*types.People, error)
}

type DBStore struct {
	db *sql.DB
}

func NewDBStore() (*DBStore, error) {
  passwordFile := os.Getenv("POSTGRES_PASSWORD_FILE")
  dbPort := os.Getenv("POSTGRES_PORT")
  dbName := os.Getenv("POSTGRES_DB")
  dbUser := os.Getenv("POSTGRES_USER")

  bin, err := ioutil.ReadFile(passwordFile)
  if err != nil {
    return nil, err
  }
  password := strings.TrimRight(string(bin), "\n")
  db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", dbUser, password, dbPort, dbName))
  if err != nil {
  	return nil, err
  }

  if err := db.Ping(); err != nil {
  	return nil, err
  }
  
  return &DBStore{
  	db: db,
  }, nil
}

func (s *DBStore) Init() error {
	return s.createTables()
}

func (s *DBStore) createTables() error {
	content, err := os.ReadFile("./schema.sql")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = s.db.Exec(string(content))
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}

func (s *DBStore) Seed() error {
	return s.seedPeople()
}

func (s *DBStore) seedPeople() error {
	for i := 0; i < 5; i++ {
		err := s.seedPerson(fmt.Sprintf("Person #%d", i))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (s *DBStore) seedPerson(name string) error {
	person, err := types.NewPerson(name)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := s.CreatePerson(person)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("new person => ", id)

	return nil
}
