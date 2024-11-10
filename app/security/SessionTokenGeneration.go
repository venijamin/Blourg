package security

import (
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
)

func GenerateSessionToken() (string, string) {
	argon := argon2.DefaultConfig()
	sessionToken := uuid.New().String()
	encoded, err := argon.HashEncoded([]byte(sessionToken))
	if err != nil {
		panic(err)
	}

	return sessionToken, string(encoded)
}
