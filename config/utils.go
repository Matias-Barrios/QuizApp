package config

import (
	"math/rand"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString :
func RandomString(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}
