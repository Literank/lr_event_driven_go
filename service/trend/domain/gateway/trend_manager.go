/*
Package gateway contains all domain gateways.
*/
package gateway

import (
	"context"

	"literank.com/event-books/domain/model"
)

// TrendManager manages all trends
type TrendManager interface {
	CreateTrend(ctx context.Context, t *model.Trend) (uint, error)
	TopTrends(ctx context.Context, pageSize uint) ([]*model.Trend, error)
}
