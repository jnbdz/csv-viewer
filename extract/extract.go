// Package extract
package extract

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// CSVFile reads CSV input from a file path
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

// CSVStdin reads CSV input from stdin
// ReadCSV reads CSV input from any io.Reader
func CSVStdin(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
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
			record[i] = strings.ReplaceAll(field, "\n", " ")
			record[i] = strings.ReplaceAll(record[i], "\r", " ")
		}

		records = append(records, record)
	}

	return records, nil
}
