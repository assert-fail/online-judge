package database

import (
	"backend/config"
	"backend/internal/pkg/logger"
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdbInstance *redis.Client
	onceRedis   sync.Once
)

func InitRedis(cfg config.Redis) (*redis.Client, error) {
	onceRedis.Do(func() {
		rdbInstance = redis.NewClient(&redis.Options{
			Addr:         cfg.Host,
			Password:     cfg.Password,
			DB:           cfg.DB,
			PoolSize:     cfg.PoolSize,
			MaxIdleConns: cfg.MaxIdleConns,
			DialTimeout:  time.Duration(cfg.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		})

		ctx := context.Background()
		pong, err := rdbInstance.Ping(ctx).Result()
		if err != nil {
			logger.Fatal().Err(err).Msg("❌ Failed to connect to Redis")
		}

		logger.Info().Msgf("Redis connected: %s", pong)
	})

	return rdbInstance, nil
}
