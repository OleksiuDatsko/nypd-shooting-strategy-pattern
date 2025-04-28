package strategy

import (
	"fmt"
	"strings"
)

type RedisOutputStrategy struct{}

func NewRedisOutputStrategy() *RedisOutputStrategy {
	return &RedisOutputStrategy{}
}

func (s *RedisOutputStrategy) Write(record []string) error {
	fmt.Println("redis", strings.Join(record, ", "))
	return nil
}

func (s *RedisOutputStrategy) Close() error {
	return nil
}
