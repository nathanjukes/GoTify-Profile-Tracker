package database

import (
	"gotify-profile-tracker/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDB struct{}

func GetPostgresDB() Database {
	return &postgresDB{}
}

func (*postgresDB) GetDB() *gorm.DB {
	connStr := os.Getenv("dbString")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal(err)
	}

	MigrateDB(db)
	return db
}

func MigrateDB(db *gorm.DB) {
	// Migrate schema
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Album{})
	db.AutoMigrate(&entity.Artist{})
	db.AutoMigrate(&entity.Track{})
	db.AutoMigrate(&entity.Activity{})
}
