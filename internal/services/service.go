package services

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type Consumer interface {
	ReadKafkaMessages(topicName string, message chan<- *sarama.ConsumerMessage)
}
type Producer interface {
	ReturnTheMessage(topicName, requestID string, data []byte) error
}

type Service struct {
	Producer Producer
	Customer Consumer
}

func NewService(producer Producer, customer Consumer) *Service {
	return &Service{
		Producer: producer,
		Customer: customer,
	}
}

// addDataToJSON takes raw data in JSON format and adds a new field "newData".
// It returns the updated data in JSON format or an error if something goes wrong.
func addDataToJSON(data []byte) ([]byte, error) {
	var receivedMessage map[string]interface{}
	if err := json.Unmarshal(data, &receivedMessage); err != nil {
		logrus.Errorf("Error unmarshaling JSON: %v\n", err)
		return nil, err
	}

	logrus.Infof("Received message: %+v\n", receivedMessage)

	receivedMessage["newData"] = "something test"
	updatedMessage, err := json.Marshal(receivedMessage)
	if err != nil {
		logrus.Errorf("Error marshaling JSON: %v\n", err)
		return nil, err
	}
	return updatedMessage, nil
}

func worker(jobs <-chan *sarama.ConsumerMessage, results chan<- *sarama.ConsumerMessage) {
	var err error
	for task := range jobs {
		// Обработка задачи
		task.Value, err = addDataToJSON(task.Value)
		if err != nil {
			logrus.Errorf("error add data to json: %v", err)
		}
		results <- task
	}
}

// GetChangeReturnJSON sets up a worker pool to process incoming Kafka messages asynchronously.
// It reads messages from the specified topic and processes them using worker goroutines.
// The updated messages are then sent back to Kafka.
func (s *Service) GetChangeReturnJSON(topicName string) {
	numWorkers := 10
	message := make(chan *sarama.ConsumerMessage, numWorkers)
	results := make(chan *sarama.ConsumerMessage, numWorkers)

	// запускаем асинхронно чтение сообщений из Kafka
	go s.Customer.ReadKafkaMessages(topicName, message)

	// Создание worker pool
	for i := 0; i < numWorkers; i++ {
		go worker(message, results)
	}

	// Обработка результатов
	for updatedMessage := range results {
		newTopicName := string(updatedMessage.Headers[0].Value)
		requestID := string(updatedMessage.Key)
		if err := s.Producer.ReturnTheMessage(newTopicName, requestID, updatedMessage.Value); err != nil {
			logrus.Errorf("Failed to send message to Kafka: %v", err)
		}
	}
}
