package utils

import (
	"os"
)

func GetEnv(key string) (value string) {

	value = os.Getenv(key)
	return
}
