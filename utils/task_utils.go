package utils

import "math/rand"

var STATUS_TYPES = [3]string{"open", "in progress", "done"}

func GetRandomStatus() string {
	index := rand.Intn(len(STATUS_TYPES))
	return STATUS_TYPES[index]
}
