package config

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
