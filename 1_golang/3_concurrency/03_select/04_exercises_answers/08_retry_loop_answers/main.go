package main

import (
	"fmt"
	"time"
)

func main() {
	success := make(chan struct{}, 1)
	attempts := 0
	for attempts < 3 {
		attempts++
		select {
		case <-success:
			fmt.Println("done")
			return
		case <-time.After(5 * time.Millisecond):
			fmt.Println("retry", attempts)
		}
	}
}
