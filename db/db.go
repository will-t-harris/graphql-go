package db

import (
	"graphql-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://pg:password@localhost:5432/graphql-go"

	db, error := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if error != nil {
		log.Fatalln(error)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
