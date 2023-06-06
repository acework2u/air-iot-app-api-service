package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("MONGURI")
}
