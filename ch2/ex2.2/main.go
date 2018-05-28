package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ruiznato/gopl/ch2/ex2.2/conv"
)

func main() {
	if len(os.Args) > 1 {
		for _, value := range os.Args[1:] {
			convert(value)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			convert(scan.Text())
		}
	}

}

func convert(value string) {
	q := conv.NewQuantity(value)
	unit := strings.ToLower(strings.Split(value, " ")[1])
	switch unit {
	case "m":
		fmt.Println(q.ToFeets())
	case "ft":
		fmt.Println(q.ToMeters())
	case "c":
		fmt.Println(q.ToFahrenheit())
	case "f":
		fmt.Println(q.ToCelsius())
	}
}
