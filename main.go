package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("exited")
	for {
		time.Sleep(1 * time.Second)
	}
}
