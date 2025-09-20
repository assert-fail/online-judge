package database

import (
	"backend/config"
	"backend/internal/pkg/logger"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis(cfg config.Redis) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MaxIdleConns: cfg.MaxIdleConns,
		DialTimeout:  time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	logger.Info().Msgf("Redis connected: %s", pong)

	return rdb, nil
}
