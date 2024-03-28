package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Path string

func Config(key string) string {
	if Path == "" { // Todo: change this shits
		Path = ".env"
	}
	err := godotenv.Load(Path)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
