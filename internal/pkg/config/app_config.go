package config

type AppConfig struct {
	KafkaBrokers []string
}

func NewAppConfig() *AppConfig {
	return new(AppConfig)
	
}
