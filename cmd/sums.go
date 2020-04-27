package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/luhring/crunch/crunch"
)

var sumsCmd = &cobra.Command{
	Use: "sums",
	Short: "Calculate sums of transaction values on a per-category basis",
	Long: "",
	Args: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		file, err := crunch.InputFile(inputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		rs, err := crunch.NewRecordSetFromFile(file)
		if err != nil {
			log.Fatal(err)
		}

		s := crunch.NewSummer(rs)
		results, err := s.Sum(columnNameCategoryForSums, columnNameAmount)
		if err != nil {
			log.Fatal(err)
		}

		for _, result := range results {
			fmt.Printf("%s: %s\n", result.Category, result.Currency)
		}
	},
}

var columnNameCategoryForSums string
var columnNameAmount string

const flagNameCategory = "category-column"
const flagNameAmount = "amount-column"

func init() {
	sumsCmd.Flags().StringVarP(&columnNameCategoryForSums, flagNameCategory, "c", "Category", "")
	sumsCmd.Flags().StringVarP(&columnNameAmount, flagNameAmount, "a", "Amount", "")
}
