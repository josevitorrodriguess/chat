package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	godotenv.Load()
	
	connStr := os.Getenv("CONNECT_STRING")
	fmt.Println("Connection string:", connStr)

	log.Println("trying to connect to database")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("ERROR: fail to connect to database")
		panic(err)
	}

	log.Println("connection sucessfully")
	return db
}

func RunMigrations(db *gorm.DB) {

	log.Println("running migrations")
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Message{})

}
