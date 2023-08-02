package utils

import (
	"context"
	"sync"
	"time"
	//"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)



var (
	once  sync.Once
	redisAddr = "10.31.117.157:6379"
	password = "ybfABC872_(&"
	rc *redis.Client
)


func NewRedisClient() (*redis.Client) {
	once.Do(func() {
		rc = new(redis.Client)
		rc := redis.NewClient(&redis.Options{
			Addr: redisAddr,
			Password: password,
			DB: 0,
			IdleTimeout: 350,
			PoolSize: 50,
		})
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		_, err := rc.Ping(ctx).Result()
		if err != nil {
			log.Fatal("Unable to connect to Redis", err)
		}
		log.Info("Connected to Redis server")
	})
	return rc
}


