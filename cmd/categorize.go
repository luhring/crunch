package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/luhring/crunch/crunch"
)

var categorizeCmd = &cobra.Command{
	Use:   "categorize",
	Short: "Determine the categorization of a transaction based on the value of a column",
	Long:  "",
	Args:  func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		classifier, err := crunch.NewClassifier(classifierFilePath)
		if err != nil {
			log.Fatal(err)
		}

		file, err := crunch.InputFile(inputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		rs, err := crunch.NewRecordSetFromFile(file)
		if err != nil {
			log.Fatal(err)
		}

		rs.ColumnNames = append(rs.ColumnNames, columnNameCategory)

		for _, record := range rs.Records {
			description, err := record.ColumnValue(columnNameClassifierInput)
			if err != nil {
				log.Fatal(err)
			}

			c, err := classifier.Classify(*description)
			if err != nil {
				c = defaultClass
			}

			record.SetColumnValue(columnNameCategory, c)
		}

		b, err := rs.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(b))
	},
}

var columnNameCategory string
var columnNameClassifierInput string
var classifierFilePath string
var defaultClass string

const flagNameClassifierInput = "column"
const flagNameClass = "class-output-column"
const flagNameFilePath = "categories-file"
const flagNameDefaultCategory = "default-category"

func init() {
	categorizeCmd.Flags().StringVarP(&columnNameClassifierInput, flagNameClassifierInput, "c", "", "name of column to categorize")
	//noinspection GoUnhandledErrorResult
	categorizeCmd.MarkFlagRequired(flagNameClassifierInput)

	categorizeCmd.Flags().StringVarP(&columnNameCategory, flagNameClass, "t", "Category", "name of new column to add with category values")

	categorizeCmd.Flags().StringVarP(&classifierFilePath, flagNameFilePath, "f", "", "path to categories file")
	//noinspection GoUnhandledErrorResult
	categorizeCmd.MarkFlagRequired(flagNameFilePath)

	categorizeCmd.Flags().StringVarP(&defaultClass, flagNameDefaultCategory, "d", "(not categorized)", "default category name")
}
