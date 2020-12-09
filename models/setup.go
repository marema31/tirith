package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/marema31/tirith/models/videogame"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "user=games password=94m35G@mes dbname=games host=palantir port=5432 sslmode=disable TimeZone=Europe/Paris"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database

	return videogame.AutoMigrate(database)
}
