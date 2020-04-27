package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/luhring/crunch/crunch"
)

var filterCmd = &cobra.Command{
	Use: "filter",
	Short: "Filter rows by a specified criterion",
	Long: "",
	Args: func(cmd *cobra.Command, args []string) error {
		if filterSpecifiedPositive && filterSpecifiedNegative {
			return errors.New("cannot filter by both positive and negative values at the same time")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file, err := crunch.InputFile(inputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		//noinspection GoUnhandledErrorResult
		defer file.Close()

		rs, err := crunch.NewRecordSetFromFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var keptRecords []crunch.Record

		for _, record := range rs.Records {
			v, err := record.ColumnValue(columnName)
			if err != nil {
				log.Fatal(err)
			}

			currency, err := crunch.NewCurrency(*v)
			if err != nil {
				log.Fatal(err)
			}

			zero := crunch.NewCurrencyZero()

			if filterSpecifiedPositive && currency.GreaterThan(zero) || filterSpecifiedNegative && currency.LessThan(zero) {
				keptRecords = append(keptRecords, record)
			}
		}

		results := crunch.NewRecordSet(rs.ColumnNames, keptRecords)

		b, err := results.Marshal()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(b))
	},
}

var columnName string
var filterSpecifiedPositive bool
var filterSpecifiedNegative bool

const flagColumnName = "column"
const flagFilterPositive = "positive"
const flagFilterNegative = "negative"

func init() {
	filterCmd.Flags().StringVarP(&columnName, flagColumnName, "c", "", "column to which to apply filter")
	//noinspection GoUnhandledErrorResult
	filterCmd.MarkFlagRequired(flagColumnName)

	filterCmd.Flags().BoolVar(&filterSpecifiedPositive, flagFilterPositive, false, "filter out rows with values in the specified column that aren't positive")
	filterCmd.Flags().BoolVar(&filterSpecifiedNegative, flagFilterNegative, false, "filter out rows with values in the specified column that aren't negative")
}
