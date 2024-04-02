/*
Package gateway contains all domain gateways.
*/
package gateway

import (
	"context"

	"literank.com/event-books/domain/model"
)

// BookManager manages all books
type BookManager interface {
	CreateBook(ctx context.Context, b *model.Book) (uint, error)
}
