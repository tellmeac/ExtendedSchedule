package config

// Config is a global configuration object.
type Config struct {
	Debug           bool   `env:"DEBUG" env-default:"false"`
	ListenAddress   string `env:"LISTEN_ADDRESS" env-default:"0.0.0.0:80"`
	BaseScheduleUrl string `env:"BASE_SCHEDULE_URL" env-default:"https://intime.tsu.ru/api/web/v1"`
	Redis           struct {
		Address string `env:"REDIS_ADDRESS" env-default:"localhost:6379"`
	}
	Database struct {
		Address string `env:"DB_ADDRESS"`
	}
}
