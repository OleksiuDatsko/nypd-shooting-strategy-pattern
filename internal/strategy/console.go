package strategy

import (
	"fmt"
	"strings"
)

type ConsoleOutputStrategy struct{}

func NewConsoleOutputStrategy() *ConsoleOutputStrategy {
	return &ConsoleOutputStrategy{}
}

func (s *ConsoleOutputStrategy) Write(record []string) error {
	fmt.Println(strings.Join(record, ", "))
	return nil
}

func (s *ConsoleOutputStrategy) Close() error {
	return nil
}
