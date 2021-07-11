package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client      *redis.Client
	hashSetName string
	ctx         context.Context
}

func (cache *Cache) Init() {
	L.Info("Cache inited")
	env := GetEnvVars()
	rdb := redis.NewClient(&redis.Options{
		Addr: env.RedisHost + ":" + env.RedisPort,
	})
	cache.client = rdb
	cache.hashSetName = "access_inf"
	cache.ctx = context.Background()
}

func (cache *Cache) setAccessInf(key string, value string) {
	cache.client.HSet(cache.ctx, cache.hashSetName, key, value)
}

func (cache *Cache) getAccessInf(key string) string {
	val := cache.client.HGet(cache.ctx, cache.hashSetName, key)
	return val.Val()
}

func (cache *Cache) shutdown(ctx context.Context) {
	cache.client.Close()
	L.Info("Cache client disconnected")
}
