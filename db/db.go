package db

import (
	"Belajar-Golang-7/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "757575"
	dbname   = "db_buku"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(entity.BookEntity{})
}

func GetDB() *gorm.DB {
	return db
}
