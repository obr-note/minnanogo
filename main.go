package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	close(queue)
	wg.Wait()
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue
		if more {
			fmt.Println("fetching", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}
