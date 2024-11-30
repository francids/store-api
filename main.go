package main

import (
	"log"
	"store-api/config"
	"store-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run()
}
