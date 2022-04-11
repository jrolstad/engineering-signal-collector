package config

import (
	"os"
	"strings"
)

type AppConfig struct {
	KafkaBrokers []string
}

func NewAppConfig() *AppConfig {
	brokersRaw := os.Getenv("kafka_endpoints")
	brokers := strings.Split(brokersRaw, ",")

	return &AppConfig{
		KafkaBrokers: brokers,
	}
}
