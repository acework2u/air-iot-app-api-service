package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DBUri      string `mapstructure:"MONGODB_LOCAL_URL"`
	Port       string `mapstructure:"PORT"`
	Origin     string `mapstructure:"CLIENT_ORIGIN"`
	UserPoolId string `mapstructure:"USER_POOL_ID"`
}

func LoadCongig(path string) (config Config, err error) {

	appMode := os.Getenv("APP_MODE")

	viper.AddConfigPath(path)

	if appMode == "dev" || appMode == "Dev" {
		viper.SetConfigType("env")
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigType("env")
		viper.SetConfigName("app")
	}

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
