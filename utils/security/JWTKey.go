package security

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetJWTKey() []byte {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtKeyString := os.Getenv("JWTKey")
	jwtKey := []byte(jwtKeyString)
	return jwtKey
}
