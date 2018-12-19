package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// init - connect to db using env vars
// auto migrate
func init() {
	// get dem env vars
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password,
	) //Build connection string

	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI) // open connection to pgdb

	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

// GetDB - Get the connected instance of the db
func GetDB() *gorm.DB {
	return db
}
