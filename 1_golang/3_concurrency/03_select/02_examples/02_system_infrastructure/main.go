package main

import (
	"fmt"
	"time"
)

func main() {
	health := make(chan string)
	go func() {
		time.Sleep(5 * time.Millisecond)
		health <- "ok"
	}()
	select {
	case msg := <-health:
		fmt.Println("health:", msg)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("health timeout")
	}
}
