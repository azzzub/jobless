package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=jobless port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Opening the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
