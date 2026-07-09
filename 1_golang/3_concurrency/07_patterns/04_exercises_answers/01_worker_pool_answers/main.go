package main

import (
	"fmt"
	"sync"
)

func worker(j <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range j {
		fmt.Println(v)
	}
}
func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}
	for i := 1; i <= 4; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
