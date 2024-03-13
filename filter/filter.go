// Package filter
package filter

import (
	"strconv"
	"strings"
)

// Column
func Columns(data [][]string, columns string) ([][]string, error) {
	selectedIndexes := strings.Split(columns, ",")
	var result [][]string

	for _, row := range data {
		var filteredRow []string
		for _, index := range selectedIndexes {
			i, err := strconv.Atoi(index)
			if err == nil {
				if i-1 < len(row) && i > 0 {
					filteredRow = append(filteredRow, row[i-1])
				}
			} else {
				return nil, err
			}
		}
		result = append(result, filteredRow)
	}

	return result, nil
}
