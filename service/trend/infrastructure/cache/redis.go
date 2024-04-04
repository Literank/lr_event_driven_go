/*
Package cache has all cache-related implementations.
*/
package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"literank.com/event-books/domain/model"
)

const (
	trendsKey      = "trends"
	queryKeyPrefix = "q-"
)

// RedisCache implements cache with redis
type RedisCache struct {
	c redis.UniversalClient
}

// NewRedisCache constructs a new RedisCache
func NewRedisCache(address, password string, db int) *RedisCache {
	r := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &RedisCache{
		c: r,
	}
}

func (r *RedisCache) CreateTrend(ctx context.Context, t *model.Trend) (uint, error) {
	// Store the search query in a sorted set
	member := t.Query
	_, err := r.c.ZScore(ctx, trendsKey, member).Result()
	if err != nil {
		if err == redis.Nil {
			// Member doesn't exist, add it with initial score of 1
			err = r.c.ZAdd(ctx, trendsKey, redis.Z{Score: 1, Member: member}).Err()
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}
	score, err := r.c.ZIncrBy(ctx, trendsKey, 1, member).Result()
	if err != nil {
		return 0, err
	}
	// Store the search query results
	k := queryKeyPrefix + t.Query
	results, err := json.Marshal(t.Books)
	if err != nil {
		return 0, err
	}
	_, err = r.c.Set(ctx, k, string(results), -1).Result()
	if err != nil {
		return 0, err
	}
	return uint(score), nil
}

func (r *RedisCache) TopTrends(ctx context.Context, pageSize uint) ([]*model.Trend, error) {
	topItems, err := r.c.ZRevRangeWithScores(ctx, trendsKey, 0, int64(pageSize)-1).Result()
	if err != nil {
		return nil, err
	}
	trends := make([]*model.Trend, 0)
	for _, item := range topItems {
		query, ok := item.Member.(string)
		if !ok {
			return nil, fmt.Errorf("invalid non-string member: %s", item.Member)
		}
		t := &model.Trend{
			Query: query,
		}
		k := queryKeyPrefix + query
		value, err := r.c.Get(ctx, k).Result()
		if err != nil {
			if err == redis.Nil {
				t.Books = make([]model.Book, 0)
				trends = append(trends, t)
				continue
			}
			return nil, err
		} else {
			if err := json.Unmarshal([]byte(value), &t.Books); err != nil {
				return nil, err
			}
		}
		trends = append(trends, t)
	}
	return trends, nil
}
