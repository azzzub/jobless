package config

import (
	"fmt"

	"github.com/azzzub/jobless/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		utils.GetEnv("PG_HOST", "localhost"),
		utils.GetEnv("PG_USER", "postgres"),
		utils.GetEnv("PG_PASS", "admin"),
		utils.GetEnv("PG_DB", "jobless"),
		utils.GetEnv("PG_PORT", "5432"),
	)

	// Opening the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
