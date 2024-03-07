package configs

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type Configs struct {
	EnvTopicName     string `env:"TOPIC_NAME"`
	EnvLogsLevel     string `env:"LOG_LEVEL"`
	EnvBrokerAddress string `env:"BROKER_ADDRESS"`
}

func NewConfig() *Configs {
	var cfg Configs
	flag.StringVar(&cfg.EnvTopicName, "t", "Handler", "The name of the topic being listened to")
	flag.StringVar(&cfg.EnvLogsLevel, "l", "info", "Set logging level")
	flag.StringVar(&cfg.EnvBrokerAddress, "b", "localhost:29092", "HTTP broker address")
	flag.Parse()
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
