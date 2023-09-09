package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FianGumilar/e-wallet/config"
	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/redis/go-redis/v9"
)

type redisCaheRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(conf *config.AppConfig) (interfaces.CacheRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Pass,
		DB:       0,
	})

	// Test connection redis using PING
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("ping error: %v", err)
	}
	fmt.Printf("Connected to redis successfully: %s", pong)

	return &redisCaheRepository{
		rdb: rdb,
	}, nil
}

// Set implements interfaces.CacheRepository.
func (r redisCaheRepository) Set(key string, entry []byte) error {
	return r.rdb.Set(context.Background(), key, entry, 15*time.Minute).Err()
}

// Get implements interfaces.CacheRepository.
func (r redisCaheRepository) Get(key string) ([]byte, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Failed get token: %s", err)
		return nil, err
	}
	log.Println("Successfully retrieved token")
	return []byte(val), nil
}
