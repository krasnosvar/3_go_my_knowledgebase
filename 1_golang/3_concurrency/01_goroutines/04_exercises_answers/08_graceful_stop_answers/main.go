package main

import (
	"fmt"
	"sync"
)

func main() {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-stop
			fmt.Println("worker", id, "stopped")
		}(i)
	}
	close(stop)
	wg.Wait()
}
