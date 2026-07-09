package main

import (
	"fmt"
	"sync"
)

func main() {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { defer wg.Done(); <-stop; fmt.Println("stopped") }()
	close(stop)
	wg.Wait()
}
