package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crunch",
	Short: "A simple utility for processing financial data from CSV files",
	Long: "",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var inputFilePath string

const flagNameInputFilePath = "input"

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFilePath, flagNameInputFilePath, "i", "", "path to the input file")
	rootCmd.AddCommand(sumsCmd)
	rootCmd.AddCommand(categorizeCmd)
	rootCmd.AddCommand(filterCmd)
}
