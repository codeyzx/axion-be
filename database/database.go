package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := "postgres://xntvhptn:vaA_cLu0kIN-75ysGKCkiT4241DhpdkD@trumpet.db.elephantsql.com/xntvhptn"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to database")
}
