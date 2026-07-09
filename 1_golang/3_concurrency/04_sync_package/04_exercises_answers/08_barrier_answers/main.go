package main

import (
	"fmt"
	"sync"
)

func main() {
	var start sync.WaitGroup
	start.Add(1)
	var workers sync.WaitGroup
	for i := 0; i < 2; i++ {
		workers.Add(1)
		go func(id int) {
			defer workers.Done()
			start.Wait()
			fmt.Println("worker", id, "started")
		}(i)
	}
	start.Done()
	workers.Wait()
}
