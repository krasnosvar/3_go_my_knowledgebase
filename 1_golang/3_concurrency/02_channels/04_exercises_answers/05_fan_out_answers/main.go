package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "job", j)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	for j := 1; j <= 4; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
