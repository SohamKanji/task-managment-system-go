package utils

import "math/rand"

var STATUS_TYPES = [3]string{"open", "in progress", "done"}

var PRIORITY_MAP = map[string]int64{"low": 0, "medium": 1, "high": 2}

func GetRandomStatus() string {
	index := rand.Intn(len(STATUS_TYPES))
	return STATUS_TYPES[index]
}

func GetPriorityValue(priority string) (int64, bool) {
	value, ok := PRIORITY_MAP[priority]
	return value, ok
}

func IsValidStatus(status string) bool {
	for _, s := range STATUS_TYPES {
		if s == status {
			return true
		}
	}
	return false
}
