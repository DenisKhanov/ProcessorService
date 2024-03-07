package main

import (
	"GoProgects/WorkTests/Exnode/FirstService/internal/configs"
	"GoProgects/WorkTests/Exnode/FirstService/internal/consumers"
	"GoProgects/WorkTests/Exnode/FirstService/internal/loggers"
	"GoProgects/WorkTests/Exnode/FirstService/internal/producers"
	"GoProgects/WorkTests/Exnode/FirstService/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	//Инициализируем переменные окружения или флаги CLI
	cfg := configs.NewConfig()

	loggers.RunLoggerConfig(cfg.EnvLogsLevel)
	logrus.Infof("Second service started with configuration logs level: %v", cfg.EnvLogsLevel)

	myProducer, err := producers.NewProducerKafka(cfg.EnvBrokerAddress)
	if err != nil {
		logrus.Fatalf("Failed to create producer: %v", err)
	}
	defer myProducer.Close()

	myConsumer, err := consumers.NewConsumerKafka(cfg.EnvBrokerAddress)
	if err != nil {
		logrus.Fatalf("Failed to create consumer: %v", err)
	}
	defer myConsumer.Close()

	myService := services.NewService(myProducer, myConsumer)

	myService.GetChangeReturnJSON(cfg.EnvTopicName)
}
