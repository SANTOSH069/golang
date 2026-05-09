package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func generate(url string, ch chan<- string) {
	st := time.Now()

	// Add https:// if missing
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Request error for %s: %v", url, err)
		return
	}

	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error while reading %s: %v", url, err)
		return
	}

	secs := time.Since(st).Seconds()

	ch <- fmt.Sprintf("%.2f %7d %s",
		secs,
		nbytes,
		url,
	)
}

func main() {
	st := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go generate(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf(
		"%.2f Time since started (Elapsed)\n",
		time.Since(st).Seconds(),
	)
}
