package processor

import (
	"fmt"
	"lab4/internal/reader"
	"lab4/internal/strategy"
)

type DataProcessor struct {
	outputStrategy strategy.OutputStrategy
}

func NewDataProcessor(strategy strategy.OutputStrategy) *DataProcessor {
	return &DataProcessor{
		outputStrategy: strategy,
	}
}

func (p *DataProcessor) ProcessData(inputFilePath string) error {
	records, err := reader.ReadCSVData(inputFilePath)
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	if len(records) == 0 {
		fmt.Println("Немає даних для обробки.")
		return nil
	}

	fmt.Printf("Починаємо обробку %d записів...\n", len(records))

	processedCount := 0
	errorCount := 0
	for i, record := range records {
		err := p.outputStrategy.Write(record)
		if err != nil {
			fmt.Printf("ERROR: error writing record #%d: %v\n", i+1, err)
			errorCount++
		} else {
			processedCount++
		}
	}

	fmt.Printf("Обробку завершено. Успішно записано: %d\n", processedCount)

	if errorCount > 0 {
		return fmt.Errorf("під час обробки виникло %d помилок запису", errorCount)
	}

	return nil
}
