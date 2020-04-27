package crunch

import (
	"fmt"
)

type Record struct {
	columns map[string]*string
}

func (r Record) ColumnValue(name string) (*string, error) {
	v := r.columns[name]
	if v == nil {
		return nil, fmt.Errorf("no value found for column with name: %s", name)
	}

	return v, nil
}

func (r Record) SetColumnValue(columnName, columnValue string) {
	r.columns[columnName] = &columnValue
}
