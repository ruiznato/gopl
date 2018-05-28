package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ruiznato/gopl/ch2/ex2.1/tempconv"
)

func main() {
	fmt.Println("*****")
	for _, value := range os.Args[1:] {
		userValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Cannot convert %q\n", value)
			fmt.Println("*****")
			continue
		}
		c := tempconv.Celsius(userValue)
		f := tempconv.CToF(c)
		k := tempconv.CToK(c)

		fmt.Println(c.String())
		fmt.Println(f.String())
		fmt.Println(k.String())
		fmt.Println("*****")
	}
}
