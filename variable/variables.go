package variable

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
