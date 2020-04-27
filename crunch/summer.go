package crunch

import (
	"sort"
)

type Category string

type SumRecord struct {
	Category Category
	Currency Currency
}

type Summer struct {
	rs *RecordSet
}

func NewSummer(rs *RecordSet) Summer {
	return Summer{
		rs: rs,
	}
}

func (s *Summer) Sum(categoryColumn, amountColumn string) ([]SumRecord, error) {
	sums := make(map[Category]Currency)

	for _, record := range s.rs.Records {
		catVal, err := record.ColumnValue(categoryColumn)
		if err != nil {
			return nil, err
		}

		category := Category(*catVal)
		amtVal, err := record.ColumnValue(amountColumn)
		if err != nil {
			return nil, err
		}

		currency, err := NewCurrency(*amtVal)
		if err != nil {
			return nil, err
		}

		sums[category] = sums[category].Add(currency)
	}

	var results []SumRecord

	for category, currency := range sums {
		results = append(results, SumRecord{
			Category: category,
			Currency: currency,
		})
	}

	sort.Slice(results, func (i, j int) bool {
		return results[i].Category < results[j].Category
	})

	return results, nil
}
