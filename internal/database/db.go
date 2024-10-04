package database

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	connStr := os.Getenv("CONNECT_STRING")

	log.Println("trying to connect to database")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("ERROR: fail to connect to database")
		panic(err)
	}

	log.Println("connection sucessfully")
	return db
}
