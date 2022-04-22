package config

import (
	"github.com/spf13/viper"
	"strings"
)

// MustLoad provides config or panic
func MustLoad() Config {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return config
}

// GetConfig returns loaded config struct
func GetConfig() (Config, error) {
	var config Config
	err := Load(&config)
	return config, err
}

func Load(config interface{}) error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("VS")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	return err
}
