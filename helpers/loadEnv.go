package helpers

import (
	"fmt"
	"log"
	"os"

	m "go-sftp-upload/models"

	"github.com/joho/godotenv"
)

// LoadEnv - Load credentials from Environment Variables (.env file)
func LoadEnv() m.Credentials {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := m.Credentials{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Server:   os.Getenv("SERVER"),
		Src:      fmt.Sprintf("\"%s\"", os.Getenv("SRC")),
		Dest:     os.Getenv("DEST"),
	}

	return env
}
