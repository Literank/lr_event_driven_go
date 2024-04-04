/*
Package consumer handles event-trigger style business logic.
*/
package consumer

import (
	"context"
	"encoding/json"

	"literank.com/event-books/domain/model"
	"literank.com/event-books/service/trend/domain/gateway"
)

type TrendConsumer struct {
	trendManager  gateway.TrendManager
	eventConsumer gateway.TrendEventConsumer
}

func NewTrendConsumer(t gateway.TrendManager, e gateway.TrendEventConsumer) *TrendConsumer {
	return &TrendConsumer{trendManager: t, eventConsumer: e}
}

func (c *TrendConsumer) Start(ctx context.Context) {
	c.eventConsumer.ConsumeEvents(ctx, func(key, data []byte) error {
		t := &model.Trend{
			Query: string(key),
		}
		if err := json.Unmarshal(data, &t.Books); err != nil {
			return err
		}
		_, err := c.trendManager.CreateTrend(ctx, t)
		return err
	})
}

func (c *TrendConsumer) EventConsumer() gateway.TrendEventConsumer {
	return c.eventConsumer
}
