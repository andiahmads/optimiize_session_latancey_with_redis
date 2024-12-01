package infra

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gored/commons/logger"
	"log/slog"
	"os"
	"time"
)

var DB0 *redis.Client

func NewRedisConn(db int) (*redis.Client, error) {
	address := os.Getenv("REDIS_DSN")
	password := os.Getenv("REDIS_PASSWORD")
	user := os.Getenv("REDIS_USER")
	client := redis.NewClient(&redis.Options{
		Addr:         address, //redis port
		Password:     password,
		Username:     user,
		DB:           db,
		PoolSize:     1000,
		PoolTimeout:  2 * time.Minute,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	DB0 = client

	logger.Slogger().Info("[INFRASTRUCTURE] Connected to the redis", pong, slog.Int("DATABASE", db))
	return DB0, nil
}
