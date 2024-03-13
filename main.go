package main

import (
	"github.com/jnbdz/csv-viewer/display"
	"github.com/jnbdz/csv-viewer/extract"
	"github.com/jnbdz/csv-viewer/filter"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"log"
	"os"
)

func main() {
	var viewMode string
	var columns string

	var rootCmd = &cobra.Command{
		Use:   "csv-viewer [filePath]",
		Short: "Display CSV content in various formats",
		Long:  `Display CSV content in various formats: column, table, json.`,
		Run: func(cmd *cobra.Command, args []string) {
			var csvData [][]string
			var err error

			// Check if stdin is a terminal or if it's receiving piped data
			if len(args) == 0 && term.IsTerminal(int(os.Stdin.Fd())) {
				// Stdin is a terminal and no arguments were provided, display help
				err := cmd.Help()
				if err != nil {
					log.Fatalf("Error: %v\n", err)
				}
				return
			}

			// If no file path is provided, read from stdin
			if len(args) == 0 {
				csvData, err = extract.CSVStdin(os.Stdin)
			} else {
				// File path is provided
				filePath := args[0]
				csvData, err = extract.CSVFile(filePath)
			}

			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}

			if columns != "" {
				csvData, err = filter.Columns(csvData, columns)
				if err != nil {
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

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
