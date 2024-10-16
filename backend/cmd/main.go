package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	storage "backend/pkg/storage"
	api "backend/pkg/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store, err := storage.NewDBStore()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Print("Prepare db...")
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	log.Print("Seed db...")
	if err := store.Seed(); err != nil {
		log.Fatal(err)
	}

	apiPort := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
	//webPort := fmt.Sprintf(":%s", os.Getenv("FRONTEND_PORT"))
	server := api.NewAPIServer(apiPort, store)
	server.Run()
}

