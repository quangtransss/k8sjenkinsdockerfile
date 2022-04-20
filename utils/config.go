package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver	string	`mapstructure:"DB_DRIVER"`
	DBSource	string	`mapstructure:"DB_SOURCE"`
	ServerAddess	string	`mapstructure:"SERVER_ADDRESS"`
}

/// load config file from .env and pass to environment


func LoadConfig(path string) (config Config,err error)  {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return

}