package utils

import (
	"os"
)

func IsFileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}

	return true
}
