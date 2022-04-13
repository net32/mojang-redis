package server

import (
	"time"

	"github.com/go-redis/redis/v8"
)

var redisCon *redis.Client

func redisConn() *redis.Client {
	if redisCon == nil {
		redisCon = redis.NewClient(&redis.Options{
			Addr:     GetEnv("REDIS_URL", "localhost:6379"),
			Password: GetEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})
	}

	return redisCon
}

type CacheResponse struct {
	hasCache bool
	response MojangResponse
}

func HasCache(cacheKey string) CacheResponse {
	rdb := redisConn()
	hasCache := rdb.Exists(rdb.Context(), cacheKey).Val()
	return CacheResponse{
		hasCache == 1,
		MojangResponse{
			200,
			rdb.Get(rdb.Context(), cacheKey).Val()}}
}

func SaveCache(cacheKey string, response MojangResponse) CacheResponse {
	saved := false
	if response.Code == 200 {
		rdb := redisConn()
		rdb.Set(rdb.Context(), cacheKey, response.Json, time.Hour).Val()
		saved = true
	}
	return CacheResponse{saved, response}
}
