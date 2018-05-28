package conv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Quantity stores the quantity to be converted
type Quantity struct {
	value float64
	unit  string
}

var units = map[string]string{
	"m":  "meters",
	"ft": "feets",
}

// NewQuantity initializes a quantity instance
func NewQuantity(s string) Quantity {
	value, unit := parse(s)
	q := Quantity{value, unit}
	return q
}

func (q Quantity) String() string {
	fmt.Println(q.value)
	fmt.Println(q.unit)
	return fmt.Sprintf("%.2f %s", q.value, units[q.unit])
}

func (q Quantity) ToMeters() string {
	meters := q.value / 3.28084
	return fmt.Sprintf("%.2f %s = %.2f meters", q.value, q.unit, meters)
}

func (q Quantity) ToFeets() string {
	feets := q.value * 3.28084
	return fmt.Sprintf("%.2f %s = %.2f feets", q.value, q.unit, feets)
}

func (q Quantity) ToFahrenheit() string {
	f := q.value*9/5 + 32
	return fmt.Sprintf("%.2f C = %.2f F", q.value, f)
}

func (q Quantity) ToCelsius() string {
	c := (q.value - 32) * 5 / 9
	return fmt.Sprintf("%.2f F = %.2f C", q.value, c)
}

func (q Quantity) To(s string) string {
	switch s {
	case "m":
		return q.ToMeters()
	case "ft":
		return q.ToFeets()
	case "c":
		return q.ToCelsius()
	case "f":
		return q.ToFahrenheit()
	default:
		return "conv: bad unit"
	}
}

func parse(s string) (float64, string) {
	values := strings.Split(s, " ")
	if len(values) != 2 {
		log.Panic(fmt.Sprintf("conv: bad value %s\n", s))
	}
	val, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		log.Panic(fmt.Sprintf("conv: bad value %s\n", s))
	}
	return val, values[1]
}
