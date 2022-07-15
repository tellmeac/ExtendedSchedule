package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

// GetConfig reads config from env once.
func GetConfig() Config {
	once.Do(func() {
		var config Config
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatal().Err(err).Msg("failed to read config from environment")
		}
		instance = &config
	})
	return *instance
}
