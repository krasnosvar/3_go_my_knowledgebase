package main

import "fmt"

func main() {
	jobs := make(chan int, 3)
	done := make(chan struct{})
	go func() {
		for j := range jobs {
			fmt.Println("job", j)
		}
		close(done)
	}()
	jobs <- 1
	jobs <- 2
	close(jobs)
	<-done
}
