/*
Package gateway contains all domain gateways.
*/
package gateway

import (
	"context"

	"literank.com/event-books/domain/model"
)

type ConsumeCallback func(key, value []byte) error

// TrendManager manages all trends
type TrendManager interface {
	CreateTrend(ctx context.Context, t *model.Trend) (uint, error)
	TopTrends(ctx context.Context, pageSize uint) ([]*model.Trend, error)
}

// TrendEventConsumer consumes trend events
type TrendEventConsumer interface {
	ConsumeEvents(ctx context.Context, callback ConsumeCallback)
	Stop() error
}
