package strategy

import (
	"context"
	"fmt"
	"lab4/internal/config"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaOutputStrategy struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaOutputStrategy(cfg config.KafkaConfig) (*KafkaOutputStrategy, error) {
	if len(cfg.Brokers) == 0 {
		return nil, fmt.Errorf("не вказані брокери Kafka")
	}
	if cfg.Topic == "" {
		return nil, fmt.Errorf("не вказана тема (topic) Kafka")
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Brokers...),
		Topic:                  cfg.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		BatchSize:              100,
		BatchTimeout:           10 * time.Millisecond,
	}

	return &KafkaOutputStrategy{
		writer: w,
		topic:  cfg.Topic,
	}, nil
}

func (s *KafkaOutputStrategy) Write(record []string) error {
	key := []byte(nil)
	if len(record) > 0 {
		key = []byte(record[0])
	}
	value := []byte(strings.Join(record, ","))

	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := s.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("error writing to Kafka (topic %s): %v", s.topic, err)
	}
	fmt.Printf("Записано в Kafka: %s\n", string(value))
	return nil
}

func (s *KafkaOutputStrategy) Close() error {
	fmt.Println("Закриття з'єднання Kafka Writer...")
	if s.writer != nil {
		err := s.writer.Close()
		if err != nil {
			return fmt.Errorf("error closing Kafka Writer: %w", err)
		}
		fmt.Println("З'єднання Kafka Writer успішно закрито.")
		return nil
	}
	return nil
}
