package database

import (
	"fmt"
	"mypackage/variable"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var username = variable.GoDotEnvVariable("DB_USER")
var password = variable.GoDotEnvVariable("DB_PASS")
var hostname = variable.GoDotEnvVariable("DB_URL")
var dbName = variable.GoDotEnvVariable("DB_NAME")

var Db *gorm.DB

func dsn() (dsnn string) {
	dsnn = username + ":" + password + "@tcp(" + hostname + ")/" + dbName
	return dsnn
}

func InitDb() *gorm.DB {
	Db = ConnectDB()
	return Db
}

func ConnectDB() *gorm.DB {
	dsn := "user:user@tcp(localhost:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}

func DbMigrate() {
	var person variable.User
	InitDb().AutoMigrate(person)
	fmt.Println("Migrating database successfully")
}


