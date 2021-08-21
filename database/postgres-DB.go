package database

import (
	"gotify-profile-tracker/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDB struct{}

func GetPostgresDB() Database {
	return &postgresDB{}
}

func (*postgresDB) GetDB() *gorm.DB {
	const connStr = "host=localhost user=postgres password=root dbname=gotify port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal(err)
	}

	MigrateDB(db)
	return db
}

func MigrateDB(db *gorm.DB) {
	// Migrate schema
	db.AutoMigrate(&entity.Activity{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Album{})
	db.AutoMigrate(&entity.Artist{})
	db.AutoMigrate(&entity.Track{})
}
