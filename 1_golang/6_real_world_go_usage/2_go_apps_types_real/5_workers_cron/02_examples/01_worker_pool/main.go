package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := []int{1, 2, 3, 4}
	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Println("processed", job)
		}(j)
	}
	wg.Wait()
}
