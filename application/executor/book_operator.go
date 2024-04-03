/*
Package executor handles request-response style business logic.
*/
package executor

import (
	"context"
	"encoding/json"
	"fmt"

	"literank.com/event-books/domain/gateway"
	"literank.com/event-books/domain/model"
	"literank.com/event-books/infrastructure/mq"
)

// BookOperator handles book input/output and proxies operations to the book manager.
type BookOperator struct {
	bookManager gateway.BookManager
	mqHelper    mq.Helper
}

// NewBookOperator constructs a new BookOperator
func NewBookOperator(b gateway.BookManager, m mq.Helper) *BookOperator {
	return &BookOperator{bookManager: b, mqHelper: m}
}

// CreateBook creates a new book
func (o *BookOperator) CreateBook(ctx context.Context, b *model.Book) (*model.Book, error) {
	id, err := o.bookManager.CreateBook(ctx, b)
	if err != nil {
		return nil, err
	}
	b.ID = id
	return b, nil
}

// GetBooks gets a list of books by offset and keyword, and caches its result if needed
func (o *BookOperator) GetBooks(ctx context.Context, offset int, query string) ([]*model.Book, error) {
	books, err := o.bookManager.GetBooks(ctx, offset, query)
	if err != nil {
		return nil, err
	}
	// Send search query and its results
	if query != "" {
		jsonData, err := json.Marshal(books)
		if err != nil {
			return nil, fmt.Errorf("failed to send event due to %w", err)
		}
		o.mqHelper.SendEvent(query, jsonData)
	}
	return books, nil
}
