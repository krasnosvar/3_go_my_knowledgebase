package main

import "fmt"

func worker(jobs <-chan int, done chan<- struct{}) {
	for j := range jobs {
		fmt.Println("job", j)
	}
	done <- struct{}{}
}

func main() {
	jobs := make(chan int, 3)
	done := make(chan struct{})
	go worker(jobs, done)
	jobs <- 1
	jobs <- 2
	close(jobs)
	<-done
}
