package config

type Config struct {
	ListenAddress string
	ScheduleAPI   struct {
		BaseURL string
	}
	Database struct {
		ConnectionAddress string
	}
}
