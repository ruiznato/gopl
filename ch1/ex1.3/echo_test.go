// Package echo_test performs benchmark between concat and join
package echo_test

import (
	"strings"
	"testing"
)

var itemList = []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}

// BenchmarkConcat tests the concat method
func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, item := range itemList {
			s += sep + item
			sep = " "
		}
	}
}

// BenchmarkJoin tests the Join method
func BenchmarkJoin(b *testing.B) {
	strings.Join(itemList, " ")
}
