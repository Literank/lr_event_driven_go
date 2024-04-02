/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	"literank.com/event-books/domain/gateway"
	"literank.com/event-books/infrastructure/config"
	"literank.com/event-books/infrastructure/database"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	sqlPersistence *database.MySQLPersistence
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	db, err := database.NewMySQLPersistence(c.DB.DSN, c.App.PageSize)
	if err != nil {
		return nil, err
	}

	return &WireHelper{
		sqlPersistence: db}, nil
}

// BookManager returns an instance of BookManager
func (w *WireHelper) BookManager() gateway.BookManager {
	return w.sqlPersistence
}
