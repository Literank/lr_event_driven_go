/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	"literank.com/event-books/infrastructure/mq"
	"literank.com/event-books/service/web/domain/gateway"
	"literank.com/event-books/service/web/infrastructure/config"
	"literank.com/event-books/service/web/infrastructure/database"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	sqlPersistence *database.MySQLPersistence
	mq             *mq.KafkaQueue
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	db, err := database.NewMySQLPersistence(c.DB.DSN, c.App.PageSize)
	if err != nil {
		return nil, err
	}
	mq, err := mq.NewKafkaQueue(c.MQ.Brokers, c.MQ.Topic)
	if err != nil {
		return nil, err
	}

	return &WireHelper{
		sqlPersistence: db, mq: mq,
	}, nil
}

// BookManager returns an instance of BookManager
func (w *WireHelper) BookManager() gateway.BookManager {
	return w.sqlPersistence
}

// MessageQueueHelper returns an instance of mq helper
func (w *WireHelper) MessageQueueHelper() mq.Helper {
	return w.mq
}
