package extract

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// readCSV
func CSVFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Could not display CSV content: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	var records [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Process each field to remove newlines
		for i, field := range record {
			// Replace newline characters with a space or another character if you prefer
			record[i] = strings.ReplaceAll(field, "\n", " ")
			record[i] = strings.ReplaceAll(record[i], "\r", " ") // Also consider carriage return for Windows
		}

		records = append(records, record)
	}

	return records, nil
}

func CSVStdin() {

}
