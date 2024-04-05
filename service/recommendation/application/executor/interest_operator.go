/*
Package executor handles request-response style business logic.
*/
package executor

import (
	"context"

	"literank.com/event-books/domain/model"
	"literank.com/event-books/service/recommendation/domain/gateway"
)

// InterestOperator handles trend input/output and proxies operations to the interest manager.
type InterestOperator struct {
	interestManager gateway.InterestManager
}

// NewInterestOperator constructs a new InterestOperator
func NewInterestOperator(t gateway.InterestManager) *InterestOperator {
	return &InterestOperator{interestManager: t}
}

// TopTrends gets the top trends order by hits in descending order
func (o *InterestOperator) InterestsForUser(ctx context.Context, userID string) ([]*model.Interest, error) {
	return o.interestManager.ListInterests(ctx, userID)
}
