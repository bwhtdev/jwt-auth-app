package storage

import (
	"database/sql"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"log"
    "math/rand/v2"

	_ "github.com/lib/pq"
	"github.com/google/uuid"
	types "backend/pkg/types"
)

type Storage interface {
	CreateUser(*types.User) (uuid.UUID, error)
	GetUserByID(string) (*types.User, error)
	GetUserByUsername(string) (*types.User, error)
	//GetUsers() ([]*types.User, error)
	UpdateUser(*types.User) error
	DeleteUser(string) error

	CreateMessage(*types.Message) (uuid.UUID, error)
	//UpdateMessage(*types.Message) error
	//DeleteMessage(string) error
	GetMessage(string) (*types.Message, error)
	GetMessages() ([]*types.Message, error)
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
	err := s.seedMessages()
	if err != nil {
		return err
	}

	return s.seedUsers()
}

func (s *DBStore) seedMessages() error {
	for i := 0; i < 5; i++ {
		randNum := rand.IntN(2) % 2
		randUsername := "user1"
		if randNum == 1 {
			randUsername = "user2"
		}
		
		err := s.seedMessage(fmt.Sprintf("Msg #%d", i), randUsername)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (s *DBStore) seedMessage(text, username string) error {
	message, err := types.NewMessage(text, username)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := s.CreateMessage(message)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("new message => ", id)

	return nil
}

func (s *DBStore) seedUsers() error {
	err := s.seedUser("user1", "password123")
	if err != nil {
		return err
	}
	return s.seedUser("user2", "password123")
}

func (s *DBStore) seedUser(username, password string) error {
	user, err := types.NewUser(username, password)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := s.CreateUser(user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("new user => ", id)

	return nil
}
