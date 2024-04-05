/*
Package mq does all message queue jobs.
*/
package mq

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"

	"literank.com/event-books/domain/gateway"
)

// KafkaConsumer consumers events from the kafka queue
type KafkaConsumer struct {
	cg    sarama.ConsumerGroup
	topic string
}

// NewKafkaConsumer constructs a new KafkaConsumer
func NewKafkaConsumer(brokers []string, topic, groupID string) (*KafkaConsumer, error) {
	// Create a new consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create consumer
	fmt.Println(brokers)
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}
	return &KafkaConsumer{consumer, topic}, nil
}

func (k *KafkaConsumer) ConsumeEvents(ctx context.Context, callback gateway.ConsumeCallback) {

	consumer := Consumer{callback}
	if err := k.cg.Consume(ctx, []string{k.topic}, &consumer); err != nil {
		log.Panicf("Failed to start consuming: %v", err)
	}
}

func (k *KafkaConsumer) Stop() error {
	return k.cg.Close()
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	callback gateway.ConsumeCallback
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Started to consume events...")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			if err := c.callback(message.Key, message.Value); err != nil {
				log.Printf("Failed to handle event from [%s] key = %s, timestamp = %v, value = %s, error: %v", message.Topic, string(message.Key), message.Timestamp, string(message.Value), err)
			}
			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
