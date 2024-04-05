/*
Package consumer handles event-trigger style business logic.
*/
package consumer

import (
	"context"
	"encoding/json"
	"strings"

	topgw "literank.com/event-books/domain/gateway"
	"literank.com/event-books/domain/model"
	"literank.com/event-books/service/recommendation/domain/gateway"
)

type InterestConsumer struct {
	interestManager gateway.InterestManager
	eventConsumer   topgw.EventConsumer
}

func NewInterestConsumer(t gateway.InterestManager, e topgw.EventConsumer) *InterestConsumer {
	return &InterestConsumer{interestManager: t, eventConsumer: e}
}

func (c *InterestConsumer) Start(ctx context.Context) {
	c.eventConsumer.ConsumeEvents(ctx, func(key, data []byte) error {
		parts := strings.Split(string(key), ":")
		if len(parts) == 1 {
			// No userID, ignore it
			return nil
		}

		var books []*model.Book
		if err := json.Unmarshal(data, &books); err != nil {
			return err
		}
		userID := parts[1]
		for _, book := range books {
			i := &model.Interest{
				UserID: userID,
				Title:  book.Title,
				Author: book.Author,
			}
			if err := c.interestManager.IncreaseInterest(ctx, i); err != nil {
				return err
			}
		}
		return nil
	})
}

func (c *InterestConsumer) EventConsumer() topgw.EventConsumer {
	return c.eventConsumer
}
