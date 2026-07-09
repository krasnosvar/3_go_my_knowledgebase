package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})
	wg.Add(1)
	go func() { defer wg.Done(); <-stop }()
	close(stop)
	wg.Wait()
	fmt.Println("clean stop")
}
