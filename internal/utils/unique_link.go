package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerateUniqueLink(email string) string {
	randomString := generateRandomString(16)
	encodedEmail := encodeEmail(email, randomString)
	return fmt.Sprintf("%sunique-link/%s", os.Getenv("BASE_URL"), encodedEmail)
}

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(randomString)
}

func encodeEmail(email, randomString string) string {
	hash := sha256.New()
	hash.Write([]byte(email + randomString))
	hashInBytes := hash.Sum(nil)
	encodedEmail := hex.EncodeToString(hashInBytes)
	return encodedEmail
}
