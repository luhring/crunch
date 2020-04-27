package crunch

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type RecordSet struct {
	ColumnNames []string
	Records []Record
}

func NewRecordSet(columnNames []string, records []Record) *RecordSet {
	return &RecordSet{
		ColumnNames: columnNames,
		Records:     records,
	}
}

func NewRecordSetFromFile(file *os.File) (*RecordSet, error) {
	if file == nil {
		return nil, errors.New("cannot create record set from file using nil file reference")
	}

	r := csv.NewReader(file)

	// identify column names
	columnNames, err := r.Read()
	if err != nil {
		return nil, err
	}

	// add records
	csvRecords, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []Record

	for _, r := range csvRecords {
		currentRecord := Record{
			columns: make(map[string]*string),
		}

		for columnIndex, columnName := range columnNames {
			currentRecord.columns[columnName] = &r[columnIndex]
		}
		records = append(records, currentRecord)
	}

	return &RecordSet{
		ColumnNames: columnNames,
		Records: records,
	}, nil
}

func (rs RecordSet) Marshal() ([]byte, error) {
	var out []byte

	headerOutput := strings.Join(rs.ColumnNames, ",")
	out = append(out, []byte(headerOutput + "\n")...)

	for _, record := range rs.Records {
		var values []string

		for _, col := range rs.ColumnNames {
			v, err := record.ColumnValue(col)
			if err != nil {
				return nil, err
			}

			values = append(values, fmt.Sprintf("\"%s\"", *v))
		}

		recordOutput := strings.Join(values, ",")
		out = append(out, []byte(recordOutput + "\n")...)
	}

	out = append(out, []byte("\n")...)

	return out, nil
}
