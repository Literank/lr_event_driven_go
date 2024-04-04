/*
Package executor handles request-response style business logic.
*/
package executor

import (
	"context"

	"literank.com/event-books/domain/model"
	"literank.com/event-books/service/trend/domain/gateway"
)

// TrendOperator handles trend input/output and proxies operations to the trend manager.
type TrendOperator struct {
	trendManager gateway.TrendManager
}

// NewTrendOperator constructs a new TrendOperator
func NewTrendOperator(t gateway.TrendManager) *TrendOperator {
	return &TrendOperator{trendManager: t}
}

// TopTrends gets the top trends order by hits in descending order
func (o *TrendOperator) TopTrends(ctx context.Context, pageSize uint) ([]*model.Trend, error) {
	return o.trendManager.TopTrends(ctx, pageSize)
}
