// Package display contains all the different ways you can view a CSV content.
//
// - Column Is the
package display

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

// Table
func Table(csvData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(csvData[0])
	table.SetAutoFormatHeaders(false) // Keep headers as they are
	table.AppendBulk(csvData[1:])     // Append data excluding header
	table.Render()
}

// JSON
func JSON(csvData [][]string) error {
	jsonData, err := json.MarshalIndent(csvData, "", "    ")
	if err != nil {
		return fmt.Errorf("Error generating JSON: %v\n", err)
	}
	_, err = fmt.Println(string(jsonData))
	return err
}

// Column view is for displaying the CSV content
func Column(csvData [][]string) error {
	var err error
	maxWidths := getColumnMaxWidths(csvData)
	for _, row := range csvData {
		for i, field := range row {
			// Adjust the padding according to the maximum width of the column
			// +2 for padding between columns
			if _, err = fmt.Printf("%-*s", maxWidths[i]+2, field); err != nil {
				return err
			}
		}
		// Newline at the end of each row
		if _, err = fmt.Println(); err != nil {
			return err
		}
	}
	return nil
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
