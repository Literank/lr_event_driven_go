/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	topgw "literank.com/event-books/domain/gateway"
	"literank.com/event-books/infrastructure/mq"
	"literank.com/event-books/service/recommendation/domain/gateway"
	"literank.com/event-books/service/recommendation/infrastructure/config"
	"literank.com/event-books/service/recommendation/infrastructure/database"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	noSQLPersistence *database.MongoPersistence
	consumer         *mq.KafkaConsumer
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	mdb, err := database.NewMongoPersistence(c.DB.MongoURI, c.DB.MongoDBName, c.App.PageSize)
	if err != nil {
		return nil, err
	}
	consumer, err := mq.NewKafkaConsumer(c.MQ.Brokers, c.MQ.Topic, c.MQ.GroupID)
	if err != nil {
		return nil, err
	}
	return &WireHelper{
		noSQLPersistence: mdb, consumer: consumer,
	}, nil
}

// InterestManager returns an instance of InterestManager
func (w *WireHelper) InterestManager() gateway.InterestManager {
	return w.noSQLPersistence
}

// TrendEventConsumer returns an instance of TrendEventConsumer
func (w *WireHelper) TrendEventConsumer() topgw.EventConsumer {
	return w.consumer
}
