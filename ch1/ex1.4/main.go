// ex1.4 prints the count and text of duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for file := range counts {
		fmt.Printf("File: %s\n", file)
		for line, n := range counts[file] {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[f.Name()] == nil {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
