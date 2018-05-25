// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(URL string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(URL)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	f, err := os.Create(url.QueryEscape(URL))
	if err != nil {
		ch <- fmt.Sprintf("whlie creating file: %v", err)
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", URL, err)
	}

	f.Close()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, URL)
}
