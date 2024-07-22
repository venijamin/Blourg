package security

import (
	"backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func OpenConnection(connectionString string) {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
}

func GetDatabase() *gorm.DB {
	return db
}
