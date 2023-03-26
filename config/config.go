package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Logging LoggingConfig
	Server  ServerConfig
}

type LoggingConfig struct {
	Level string
}

type ServerConfig struct {
	Host string
	Port string
	Mode string
}

func ReadConfig() *Config {
	viper.AddConfigPath("/config")
	viper.AddConfigPath("./config/files")

	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic("An error occurred while reading the configuration file: " + err.Error())
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("An error occurred while reading the configuration file: " + err.Error())
	}

	return &cfg
}
