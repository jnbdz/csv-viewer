package filter

import (
	"reflect"
	"testing"
)

func TestColumns(t *testing.T) {
	// Test data setup
	csvData := [][]string{
		{"Name", "Age", "Email"},
		{"John Doe", "30", "john@example.com"},
		{"Jane Doe", "25", "jane@example.com"},
	}

	// Define test cases
	tests := []struct {
		name    string
		columns string
		want    [][]string
		wantErr bool
	}{
		{
			name:    "Filter single column",
			columns: "1",
			want: [][]string{
				{"Name"},
				{"John Doe"},
				{"Jane Doe"},
			},
			wantErr: false,
		},
		{
			name:    "Filter multiple columns",
			columns: "1,3",
			want: [][]string{
				{"Name", "Email"},
				{"John Doe", "john@example.com"},
				{"Jane Doe", "jane@example.com"},
			},
			wantErr: false,
		},
		{
			name:    "Invalid column index",
			columns: "0",
			want:    nil,
			wantErr: true, // Expect an error due to zero or negative index
		},
		{
			name:    "Out of range column index",
			columns: "4",
			want: [][]string{
				{},
				{},
				{},
			},
			wantErr: false, // Out of range indexes result in missing data but not an error
		},
		{
			name:    "Invalid column index format",
			columns: "abc",
			want:    nil,
			wantErr: true, // Expect an error due to non-integer string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Columns(csvData, tt.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("Columns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Columns() got = %v, want %v", got, tt.want)
			}
		})
	}
}
