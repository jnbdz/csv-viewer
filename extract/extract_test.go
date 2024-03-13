package extract

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestCSVFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "example.csv")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	// Write CSV content to the temporary file
	_, err = tmpFile.WriteString("Name,Age\nJohn,30\nJane,25")
	if err != nil {
		t.Fatalf("Unable to write to temporary file: %v", err)
	}
	tmpFile.Close() // Close the file so it can be reopened by CSVFile

	// Call CSVFile
	records, err := CSVFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("CSVFile returned an error: %v", err)
	}

	// Expected records
	expected := [][]string{
		{"Name", "Age"},
		{"John", "30"},
		{"Jane", "25"},
	}

	if !reflect.DeepEqual(records, expected) {
		t.Errorf("CSVFile() = %v, want %v", records, expected)
	}
}

func TestReadCSV(t *testing.T) {
	content := "Name,Age\nJohn,30\nJane,25"
	reader := strings.NewReader(content)

	records, err := CSVStdin(reader)
	if err != nil {
		t.Fatalf("ReadCSV returned an error: %v", err)
	}

	expected := [][]string{
		{"Name", "Age"},
		{"John", "30"},
		{"Jane", "25"},
	}

	if !reflect.DeepEqual(records, expected) {
		t.Errorf("ReadCSV() = %v, want %v", records, expected)
	}
}
