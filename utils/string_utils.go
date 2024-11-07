package utils

import (
	"math/rand"
)

var LETTERS = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomString(length int) string {
	res := make([]byte, length)
	for i := range res {
		index := rand.Intn(len(LETTERS))
		res[i] = LETTERS[index]
	}
	return string(res)
}
