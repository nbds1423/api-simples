package utils

import (
	"fmt"
	"log"
	"os"
	"personal-api/src/structs"

	"github.com/joho/godotenv"
)

var Env structs.Env

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[Error] - .env file not found.")
	}

	Env = structs.Env{
		Port: port(),
	}
}

func port() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	port := "3000"
	fmt.Printf("\n-> Port is missing in .env! \n-> For security, your port has been set to %v\n\n", port)
	return ":" + port
}
