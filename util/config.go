package util

import "github.com/spf13/viper"

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
}
