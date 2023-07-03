package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBUri          string `mapstructure:"MONGODB_LOCAL_URL"`
	Port           string `mapstructure:"PORT"`
	Origin         string `mapstructure:"CLIENT_ORIGIN"`
	UserPoolId     string `mapstructure:"USER_POOL_ID"`
	CognRegion     string `mapstructure:"COGNITO_REGION"`
	CognClientId   string `mapstructure:"COGNITO_CLIENT_ID"`
	CognUserPoolId string `mapstructure:"COGNITO_USER_POOL_ID"`
}

func LoadCongig(path string) (config Config, err error) {

	appMode := os.Getenv("APP_MODE")

	fileEnvType := "env"
	fileEnvName := "app"

	envPath := "~/air-iot-app-api-service/bin"

	if len(path) > 0 {
		envPath = path
	}

	if appMode == "dev" || appMode == "Dev" {
		fileEnvName = "dev"
		envPath = "."
	}

	viper.AddConfigPath(envPath)
	viper.SetConfigType(fileEnvType)
	viper.SetConfigName(fileEnvName)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
