/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	"literank.com/event-books/service/trend/domain/gateway"
	"literank.com/event-books/service/trend/infrastructure/cache"
	"literank.com/event-books/service/trend/infrastructure/config"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	kvStore *cache.RedisCache
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	kv := cache.NewRedisCache(c.Cache.Address, c.Cache.Password, c.Cache.DB)
	return &WireHelper{
		kvStore: kv,
	}, nil
}

// TrendManager returns an instance of TrendManager
func (w *WireHelper) TrendManager() gateway.TrendManager {
	return w.kvStore
}
