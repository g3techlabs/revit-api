package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetIntVariable(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic("Error parsing " + key + " to int type: " + err.Error())
	}
	return value
}
