/*
Package mq does all message queue jobs.
*/
package mq

import (
	"github.com/IBM/sarama"
)

// KafkaQueue runs all Kafka operations
type KafkaQueue struct {
	producer sarama.SyncProducer
	topic    string
}

// NewKafkaQueue constructs a new KafkaQueue
func NewKafkaQueue(brokers []string, topic string) (*KafkaQueue, error) {
	// Configuration
	c := sarama.NewConfig()
	c.Producer.RequiredAcks = sarama.WaitForLocal
	c.Producer.Retry.Max = 3
	c.Producer.Return.Successes = true
	// Create producer
	producer, err := sarama.NewSyncProducer(brokers, c)
	if err != nil {
		return nil, err
	}
	return &KafkaQueue{producer, topic}, nil
}

// CreateBook creates a new book
func (k *KafkaQueue) SendEvent(key string, value []byte) (bool, error) {
	// Send a message
	message := &sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	// Send message to Kafka
	_, _, err := k.producer.SendMessage(message)
	if err != nil {
		return false, err
	}
	return true, nil
}
