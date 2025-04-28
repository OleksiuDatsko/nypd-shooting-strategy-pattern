package strategy

import (
	"fmt"
	"lab4/internal/config"
)

type OutputStrategy interface {
	Write(record []string) error
	Close() error
}

func CreateOutputStrategy(cfg *config.Config) (OutputStrategy, error) {
	switch cfg.OutputType {
	case "console":
		fmt.Println("Використовується стратегія: Console")
		return NewConsoleOutputStrategy(), nil
	case "kafka":
		fmt.Println("Використовується стратегія: Kafka")
		if len(cfg.Kafka.Brokers) == 0 || cfg.Kafka.Topic == "" {
			return nil, fmt.Errorf("for type 'kafka' required fields 'kafka.brokers' and 'kafka.topic'")
		}
		return NewKafkaOutputStrategy(), nil
	case "redis":
		fmt.Println("Використовується стратегія: Redis")
		if cfg.Redis.Address == "" || cfg.Redis.ListKey == "" {
			return nil, fmt.Errorf("for type 'redis' required fields 'redis.address' and 'redis.list_key'")
		}
		return NewRedisOutputStrategy(), nil
	default:
		return nil, fmt.Errorf("unknown output type: %s", cfg.OutputType)
	}
}
