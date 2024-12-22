package registration

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
)

// Хэширование пароля с использованием SHA-256
func hashPassword(password string) string {
	hash := sha256.New()
	io.WriteString(hash, password)
	return hex.EncodeToString(hash.Sum(nil))
}

// Генерация случайной соли
func generateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(salt)
}
