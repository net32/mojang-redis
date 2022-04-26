package server

import (
	"encoding/json"
	"log"
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
	HasCache bool           `json:"hascache"`
	Response MojangResponse `json:"response"`
}

func ExistsKey(cacheKey string) int64 {
	rdb := redisConn()
	return rdb.Exists(rdb.Context(), cacheKey).Val()
}

func SaveValue(cacheKey string, value string) bool {
	rdb := redisConn()
	rdb.Set(rdb.Context(), cacheKey, value, 24*time.Hour).Val()
	return true
}

func GetValue(cacheKey string) string {
	rdb := redisConn()
	return rdb.Get(rdb.Context(), cacheKey).Val()
}

func HasCache(cacheKey string) CacheResponse {
	hasCache := ExistsKey(cacheKey)
	value := GetValue(cacheKey)
	var response MojangResponse
	if hasCache == 1 {
		err := json.Unmarshal([]byte(value), &response)
		if err != nil {
			log.Println(err, "HasCache:", cacheKey)
		}
	}
	return CacheResponse{hasCache == 1, response}
}

func SaveCache(cacheKey string, response MojangResponse) CacheResponse {
	saved := false
	if response.Code < 500 {
		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err, "SaveCache:", cacheKey, response)
		}
		saved = SaveValue(cacheKey, string(data))
	}
	return CacheResponse{saved, response}
}
