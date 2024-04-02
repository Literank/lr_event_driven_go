/*
Package executor handles request-response style business logic.
*/
package executor

import (
	"context"

	"literank.com/event-books/domain/gateway"
	"literank.com/event-books/domain/model"
)

// BookOperator handles book input/output and proxies operations to the book manager.
type BookOperator struct {
	bookManager gateway.BookManager
}

// NewBookOperator constructs a new BookOperator
func NewBookOperator(b gateway.BookManager) *BookOperator {
	return &BookOperator{bookManager: b}
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
