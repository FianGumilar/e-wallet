package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(n int) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	charsets := []rune("abcdefghijklmABCDEFGHJKLMN1234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[r.Intn(len(charsets))]
	}
	return string(letters)
}
