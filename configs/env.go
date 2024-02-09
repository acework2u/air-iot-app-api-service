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
		fmt.Println(err.Error())
		return os.Getenv("MONGURI")
	}
	return os.Getenv("MONGURI")
}
func EnvDB2() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err.Error())
		return os.Getenv("DB2")
	}
	return os.Getenv("DB2")
}
