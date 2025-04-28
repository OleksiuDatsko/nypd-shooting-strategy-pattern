package strategy

import (
	"context"
	"fmt"
	"lab4/internal/config"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOutputStrategy struct {
	client  *redis.Client
	listKey string
}

func NewRedisOutputStrategy(cfg config.RedisConfig) (*RedisOutputStrategy, error) {
	if cfg.Address == "" {
		return nil, fmt.Errorf("address redis required")
	}
	if cfg.ListKey == "" {
		return nil, fmt.Errorf("list key redis required")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisOutputStrategy{
		client:  rdb,
		listKey: cfg.ListKey,
	}, nil
}

func (s *RedisOutputStrategy) Write(record []string) error {
	value := strings.Join(record, ",")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.client.RPush(ctx, s.listKey, value).Err()
	if err != nil {
		return fmt.Errorf("error writing to Redis (key %s): %v", s.listKey, err)
	}
	fmt.Printf("Записано в Redis list '%s': %s\n", s.listKey, value) // Логування
	return nil
}

func (s *RedisOutputStrategy) Close() error {
	fmt.Println("Закриття з'єднання Redis...")
	if s.client != nil {
		err := s.client.Close()
		if err != nil {
			return fmt.Errorf("error closing Redis: %v", err)
		}
		fmt.Println("З'єднання Redis успішно закрито.")
		return nil
	}
	return nil
}
