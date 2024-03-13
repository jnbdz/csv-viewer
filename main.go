package main

import (
	"github.com/jnbdz/csv-viewer/display"
	"github.com/jnbdz/csv-viewer/extract"
	"github.com/jnbdz/csv-viewer/filter"
	"github.com/spf13/cobra"
	"log"
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

			csvData, err := extract.CSVFile(filePath)
			if err != nil {
				log.Fatalf("Error reading CSV file: %v\n", err)
			}

			if columns != "" {
				if csvData, err = filter.Columns(csvData, columns); err != nil {
					log.Fatalf("Could not filter CSV content: %v", err)
				}
			}

			switch viewMode {
			case "table":
				display.Table(csvData)
			case "json":
				err = display.JSON(csvData)
			default:
				err = display.Column(csvData)
			}

			if err != nil {
				log.Fatalf("Could not display CSV content: %v", err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&viewMode, "view", "v", "column", "View mode: column, table, json")
	rootCmd.Flags().StringVarP(&columns, "columns", "c", "", "Select columns to display (e.g., --columns=\"1,3\")")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("There seems to be : %v", err)
	}
}
