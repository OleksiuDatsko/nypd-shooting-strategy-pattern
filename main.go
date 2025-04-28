package main

import (
	"fmt"
	"lab4/internal/config"
	"lab4/internal/processor"
	"lab4/internal/strategy"
	"os"
)

func main() {
	fmt.Println("Запуск програми обробки даних...")

	configPath := "config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("He вдалося завантажити конфігурацію з %s\n\t%v\n", configPath, err)
		os.Exit(1)
	}
	fmt.Printf("Конфігурацію завантажено: Input='%s', OutputType='%s'\n", cfg.InputFile, cfg.OutputType)

	outputStrategy, err := strategy.CreateOutputStrategy(cfg)
	if err != nil {
		fmt.Printf("He вдалося створити стратегію виводу\n\t%v\n", err)
		os.Exit(1)
	}

	defer func() {
		if err := outputStrategy.Close(); err != nil {
			fmt.Printf("Помилка під час закриття ресурсів стратегії\n%v\n", err)
		} else {
			fmt.Println("Ресурси успішно закрито")
		}
	}()

	dataProcessor := processor.NewDataProcessor(outputStrategy)

	fmt.Printf("Початок обробки файлу: %s\n", cfg.InputFile)
	err = dataProcessor.ProcessData(cfg.InputFile)
	if err != nil {
		fmt.Printf("Під час обробки даних виникла помилка\n\t%v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Обробка даних успішно завершена")
	}

	fmt.Println("Програма завершила роботу")
}
