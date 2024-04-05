/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	topgw "literank.com/event-books/domain/gateway"
	"literank.com/event-books/infrastructure/mq"
	"literank.com/event-books/service/trend/domain/gateway"
	"literank.com/event-books/service/trend/infrastructure/cache"
	"literank.com/event-books/service/trend/infrastructure/config"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	kvStore  *cache.RedisCache
	consumer *mq.KafkaConsumer
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	kv := cache.NewRedisCache(c.Cache.Address, c.Cache.Password, c.Cache.DB)
	consumer, err := mq.NewKafkaConsumer(c.MQ.Brokers, c.MQ.Topic, c.MQ.GroupID)
	if err != nil {
		return nil, err
	}
	return &WireHelper{
		kvStore:  kv,
		consumer: consumer,
	}, nil
}

// TrendManager returns an instance of TrendManager
func (w *WireHelper) TrendManager() gateway.TrendManager {
	return w.kvStore
}

// TrendEventConsumer returns an instance of TrendEventConsumer
func (w *WireHelper) TrendEventConsumer() topgw.EventConsumer {
	return w.consumer
}
