package main

import "fmt"

func main() {
	done := make(chan struct{}, 1)
	done <- struct{}{}
	for {
		select {
		case <-done:
			fmt.Println("stop")
			return
		default:
		}
	}
}
