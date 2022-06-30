package configs

type MainConfig struct {
	Database DBConfig
	Auth     AuthConfig
	Port     int32
}
