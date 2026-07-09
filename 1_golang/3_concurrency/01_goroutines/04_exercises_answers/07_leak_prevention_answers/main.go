package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(5 * time.Millisecond)
			}
		}
	}()
	time.Sleep(20 * time.Millisecond)
	close(stop)
	fmt.Println("stopped")
}
