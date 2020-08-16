package libs

import (
	"crypto/rand"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func GenerateHash(password string) ([]byte, []byte) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	dk := pbkdf2.Key([]byte(password), salt, 100000, 128, sha512.New)
	return salt, dk
}
