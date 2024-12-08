package registration

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

// Хэширование пароля с использованием MD5
func hashPassword(password string) string {
	hash := md5.New()
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
