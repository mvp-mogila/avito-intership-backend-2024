package redis

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/go-redis/redis"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/config"
	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
	st "github.com/mvp-mogila/avito-intership-backend-2024/internal/pkg/storage"
)

type RedisCache struct {
	client         *redis.Client
	expirationTime time.Duration
}

func NewRedisCache(cfg config.RedisConfig) (st.Cache, error) {
	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: hostPort,
	})
	if status := redisClient.Ping(); status.Err() != nil {
		return nil, status.Err()
	}
	log.Println("redis connection opened ...")
	return &RedisCache{
		client:         redisClient,
		expirationTime: cfg.ExpTime,
	}, nil
}

func (r *RedisCache) Set(key string, value interface{}) error {
	return r.client.Set(key, value, r.expirationTime).Err()
}

func (r *RedisCache) Get(key string) ([]byte, error) {
	res := r.client.Get(key)
	if res.Err() != nil {
		if errors.Is(res.Err(), redis.Nil) {
			return nil, models.ErrEmptyCache
		}
		return nil, res.Err()
	}

	return res.Bytes()
}

func (r *RedisCache) Close() error {
	log.Println("redis connection closing ...")
	return r.client.Close()
}
