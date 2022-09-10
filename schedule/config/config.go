package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"sync"
)

// Config is a global configuration object.
type Config struct {
	Debug           bool   `env:"DEBUG" env-default:"false"`
	ListenAddress   string `env:"LISTEN_ADDRESS" env-default:"0.0.0.0:80"`
	BaseScheduleUrl string `env:"BASE_SCHEDULE_URL" env-default:"https://intime.tsu.ru/api/web/v1"`
}

var (
	instance *Config
	once     sync.Once
)

// Get reads config from env once.
func Get() Config {
	once.Do(func() {
		var config Config
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatal().Err(err).Msg("failed to read config from environment")
		}
		instance = &config
	})
	return *instance
}
