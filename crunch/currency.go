package crunch

import (
	"fmt"
	"math"
	"strconv"
)

type Currency struct {
	value int64
}

func NewCurrency(v string) (Currency, error) {
	parsed, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return Currency{}, err
	}

	adjusted := math.Round(parsed * 100)
	value := int64(adjusted)

	return Currency{
		value: value,
	}, nil
}

func NewCurrencyZero() Currency {
	return Currency{value:0}
}

func (c Currency) Add(v Currency) Currency {
	return Currency{
		value: c.value + v.value,
	}
}

func (c Currency) String() string {
	converted := float64(c.value) / 100
	return fmt.Sprint(converted)
}

func (c Currency) LessThan(other Currency) bool {
	return c.value < other.value
}

func (c Currency) GreaterThan(other Currency) bool {
	return c.value > other.value
}
