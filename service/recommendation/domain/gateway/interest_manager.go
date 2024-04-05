/*
Package gateway contains all domain gateways.
*/
package gateway

import (
	"context"

	"literank.com/event-books/domain/model"
)

// InterestManager manages all interests
type InterestManager interface {
	IncreaseInterest(ctx context.Context, i *model.Interest) error
	ListInterests(ctx context.Context, userID string) ([]*model.Interest, error)
}
