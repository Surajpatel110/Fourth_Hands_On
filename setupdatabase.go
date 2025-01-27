package services

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// setupDatabase initializes the connection to the PostgreSQL database using GORM and sets up connection pooling
func setupDatabase() *gorm.DB {
	// Database connection string format: "postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable"
	dbConStr := "host=postgres user=postgres password=9315751488 dbname=handson4 port=5432 sslmode=disable TimeZone=UTC"
	Db, err := gorm.Open(postgres.Open(dbConStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, _ := Db.DB()

	sqlDB.SetMaxOpenConns(20) // Adjust based on system capacity
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	// Db.DropTableIfExists(&Record{})
	Db.AutoMigrate(&Record{})

	return Db
}
