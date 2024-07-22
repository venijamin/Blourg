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

func VerifySession(sessionToken, encoded string) bool {
	isMatch, err := argon2.VerifyEncoded([]byte(sessionToken), []byte(encoded))
	if err != nil {
		panic(err)
	}
	return isMatch
}
