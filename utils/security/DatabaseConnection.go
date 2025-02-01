package security

import (
	"blourg/model/Post"
	"blourg/model/User"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var mainDB *gorm.DB

func GetMainDB() *gorm.DB { return mainDB }

func ConnectToDB() {
	// Connect to database
	// Load the connection string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Main database where all data is stored
	connectionString := os.Getenv("ConnectionString")
	mainDB = openConnection(connectionString)
	migrateSchemaMainDB()
}

// Whenever you want to add a new model to the database as a table you have to specify it here
func migrateSchemaMainDB() {
	// Migrate the schema
	var err error
	err = mainDB.AutoMigrate(&Post.Post{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	err = mainDB.AutoMigrate(&User.User{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
}

func openConnection(connectionString string) *gorm.DB {
	var db *gorm.DB

	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
