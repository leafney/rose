package rand

import (
	"math/rand"
	"time"
)

// String returns a random string ['a', 'z'] in the specified length
func String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// Int returns a random integer in range [min, max].
func Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min)
}
