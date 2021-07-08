package db

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/database/postgres"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to database"

func Initdatabase(databaseURL string) {
	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		fmt.Println(dbErrorMessage)
	}

	fmt.Println(db)
}
