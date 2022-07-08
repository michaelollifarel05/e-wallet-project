package variable

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	// "gorm.io/gorm"
)


type Account struct {
	UserID int64
	Balance int 	
	// User User
}

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"username"`
	Password string `json:"passwords"`
	Account Account `json:"accounts"`
}




func GoDotEnvVariable(key string) string {

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
