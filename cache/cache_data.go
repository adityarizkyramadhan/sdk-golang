package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheService struct {
	Redis *redis.Client
	Cache *cache.Cache
}

func NewCacheService(address string, size int, TTL time.Duration) *CacheService {
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})

	cache := cache.New(&cache.Options{
		Redis: rdb,
		// Cache 10k keys for 1 minute.
		LocalCache: cache.NewTinyLFU(size, TTL),
	})
	return &CacheService{
		Redis: rdb,
		Cache: cache,
	}
}

func (c *CacheService) AddData(ctx context.Context, key string, data interface{}, TTL time.Duration) error {
	item := &cache.Item{
		Ctx: ctx,
		Key: key,
		TTL: TTL,
		Do: func(*cache.Item) (interface{}, error) {
			return data, nil
		},
	}
	return c.Cache.Set(item)
}

func (c *CacheService) GetData(ctx context.Context, key string, data interface{}) error {
	return c.Cache.Get(ctx, key, data)
}

func (c *CacheService) Close() error {
	return c.Redis.Close()
}
