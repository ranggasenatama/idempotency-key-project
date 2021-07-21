package main

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	client      *redis.Client
	ttlCache    time.Duration
	ttlProgress time.Duration
)

func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ttlCache = 86400 * time.Second // one day
	ttlProgress = 60 * time.Second // one minute
}
