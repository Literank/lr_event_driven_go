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

// GetBooks gets a list of books by offset and keyword
func (s *MySQLPersistence) GetBooks(ctx context.Context, offset int, keyword string) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	tx := s.db.WithContext(ctx)
	if keyword != "" {
		term := "%" + keyword + "%"
		tx = tx.Where("title LIKE ?", term).Or("author LIKE ?", term).Or("description LIKE ?", term)
	}
	if err := tx.Offset(offset).Limit(s.pageSize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
