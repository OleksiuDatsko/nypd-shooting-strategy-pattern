package reader

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVData(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	fmt.Printf("Успішно прочитано %d записів з файлу %s\n", len(records), filePath)
	return records, nil
}