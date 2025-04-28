package strategy

import (
	"fmt"
	"strings"
)

type KafkaOutputStrategy struct{}

func NewKafkaOutputStrategy() *KafkaOutputStrategy {
	return &KafkaOutputStrategy{}
}

func (s *KafkaOutputStrategy) Write(record []string) error {
	fmt.Println("kafka", strings.Join(record, ", "))
	return nil
}

func (s *KafkaOutputStrategy) Close() error {
	return nil
}
