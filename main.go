package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var viewMode string
	var columns string

	var rootCmd = &cobra.Command{
		Use:   "csv-viewer [filePath]",
		Short: "Display CSV content in various formats",
		Long:  `Display CSV content in various formats: column, table, json.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]

			csvData, err := readCSV(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
				os.Exit(1)
			}

			if columns != "" {
				csvData = filterColumns(csvData, columns)
			}

			switch viewMode {
			case "table":
				displayTable(csvData)
			case "json":
				displayJSON(csvData)
			default:
				displayColumn(csvData)
			}
		},
	}

	rootCmd.Flags().StringVarP(&viewMode, "view", "v", "column", "View mode: column, table, json")
	rootCmd.Flags().StringVarP(&columns, "columns", "c", "", "Select columns to display (e.g., --columns=\"1,3\")")

	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

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

func filterColumns(data [][]string, columns string) [][]string {
	selectedIndexes := strings.Split(columns, ",")
	var result [][]string

	for _, row := range data {
		var filteredRow []string
		for _, index := range selectedIndexes {
			i, err := strconv.Atoi(index)
			if err == nil && i-1 < len(row) && i > 0 {
				filteredRow = append(filteredRow, row[i-1])
			}
		}
		result = append(result, filteredRow)
	}

	return result
}

func displayTable(csvData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(csvData[0])
	table.SetAutoFormatHeaders(false) // Keep headers as they are
	table.AppendBulk(csvData[1:])     // Append data excluding header
	table.Render()
}

func displayJSON(csvData [][]string) {
	jsonData, err := json.MarshalIndent(csvData, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func displayColumn(csvData [][]string) {
	maxWidths := getColumnMaxWidths(csvData)
	for _, row := range csvData {
		for i, field := range row {
			// Adjust the padding according to the maximum width of the column
			fmt.Printf("%-*s", maxWidths[i]+2, field) // +2 for padding between columns
		}
		fmt.Println() // Newline at the end of each row
	}
}

// getColumnMaxWidths returns the maximum width for each column
func getColumnMaxWidths(csvData [][]string) []int {
	maxWidths := make([]int, len(csvData[0]))
	for _, row := range csvData {
		for i, field := range row {
			if len(field) > maxWidths[i] {
				maxWidths[i] = len(field)
			}
		}
	}
	return maxWidths
}
