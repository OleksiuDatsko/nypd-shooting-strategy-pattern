package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	ListKey  string `yaml:"list_key"`
}

type Config struct {
	InputFile  string      `yaml:"input_file"`
	OutputType string      `yaml:"output_type"`
	Kafka      KafkaConfig `yaml:"kafka"`
	Redis      RedisConfig `yaml:"redis"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s: %w", path, err)
	}

	var cfg Config
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	// Додати валідацію значень конфігурації тут, якщо потрібно
	if cfg.InputFile == "" {
		return nil, fmt.Errorf("field 'input_file' required in config")
	}
	if cfg.OutputType == "" {
		return nil, fmt.Errorf("field 'output_type' required in config")
	}

	return &cfg, nil
}
