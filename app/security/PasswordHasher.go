package security

import (
	"github.com/matthewhartstonge/argon2"
)

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		panic(err)
	}

	return string(encoded)
}

func VerifyPassword(password, encoded string) bool {
	isMatch, err := argon2.VerifyEncoded([]byte(password), []byte(encoded))
	if err != nil {
		panic(err)
	}
	return isMatch
}
