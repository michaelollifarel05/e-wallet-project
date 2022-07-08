package database

import (
	"fmt"
	"mypackage/variable"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var username = variable.GoDotEnvVariable("DB_USER")
var password = variable.GoDotEnvVariable("DB_PASS")
var hostname = variable.GoDotEnvVariable("DB_URL")
var dbName = variable.GoDotEnvVariable("DB_NAME")

var Db *gorm.DB

func dsn() (dsnn string) {
	dsnn = username + ":" + password + "@tcp(" + hostname + ")/" + dbName
	fmt.Println(dsnn)
	return dsnn
}

func InitDb() *gorm.DB {
	Db = ConnectDB()
	return Db
}

func ConnectDB() *gorm.DB {
	// dsn := "user:user@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),

	})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}

func DbMigrate() {
	type Company struct {
		ID   int
		Name string
	  }
	type User struct {
		gorm.Model
		Name         string
		CompanyRefer int
		Company      Company `gorm:"foreignKey:CompanyRefer"`
		// use CompanyRefer as foreign key
	  }
	  

	  

	InitDb().AutoMigrate(variable.User{}, variable.Account{})
	// InitDb().Model(&variable.Account{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	fmt.Println("Migrating database successfully")
}

