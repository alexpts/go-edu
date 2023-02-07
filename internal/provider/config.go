package provider

import (
	"github.com/spf13/viper"
)

type Config = viper.Viper

func ProvideConfig() *Config {
	config := viper.New()
	config.AllowEmptyEnv(true)
	config.AutomaticEnv()

	return config
}
