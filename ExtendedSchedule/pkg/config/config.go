package config

// Config is a global configuration object.
type Config struct {
	IsDebug       bool
	ListenAddress string
	ScheduleAPI   struct {
		BaseURL string
	}
	Database struct {
		ConnectionAddress string
	}
}
