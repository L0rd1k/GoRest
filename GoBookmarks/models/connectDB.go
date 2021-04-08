package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var dataBase *gorm.DB

func init() {
	configFile := godotenv.Load("./data/config.env")
	if configFile != nil {
		fmt.Println(configFile)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dataBaseUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dataBaseUri)

	connected, error := gorm.Open("postgres", dataBaseUri)
	if error != nil {
		fmt.Println("Can't connect to DataBase!!!")
		fmt.Print(error)
	} else {
		dataBase = connected
		dataBase.Debug().AutoMigrate(&Account{}, &Contact{})
	}
}

func GetDataBase() *gorm.DB {
	return dataBase
}
