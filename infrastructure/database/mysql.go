/*
Package database does all db persistence implementations.
*/
package database

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"literank.com/event-books/domain/model"
)

// MySQLPersistence runs all MySQL operations
type MySQLPersistence struct {
	db       *gorm.DB
	pageSize int
}

// NewMySQLPersistence constructs a new MySQLPersistence
func NewMySQLPersistence(dsn string, pageSize int) (*MySQLPersistence, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto Migrate the data structs
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		return nil, err
	}

	return &MySQLPersistence{db, pageSize}, nil
}

// CreateBook creates a new book
func (s *MySQLPersistence) CreateBook(ctx context.Context, b *model.Book) (uint, error) {
	if err := s.db.WithContext(ctx).Create(b).Error; err != nil {
		return 0, err
	}
	return b.ID, nil
}
