// Package display

package display

import (
	"reflect"
	"testing"
)

func TestGetColumnMaxWidths(t *testing.T) {
	csvData := [][]string{
		{"Name", "Country", "Email"},
		{"John Doe", "USA", "john@example.com"},
		{"Jane Smith", "Canada", "jane@example.ca"},
	}

	expectedWidths := []int{10, 7, 16} // Corrected to include header lengths
	actualWidths := getColumnMaxWidths(csvData)

	if !reflect.DeepEqual(actualWidths, expectedWidths) {
		t.Errorf("getColumnMaxWidths() = %v, want %v", actualWidths, expectedWidths)
	}
}
