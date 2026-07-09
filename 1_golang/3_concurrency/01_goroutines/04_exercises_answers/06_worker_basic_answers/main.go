package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan struct{})
	go func() {
		for j := range jobs {
			fmt.Println("job", j)
		}
		close(done)
	}()
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)
	<-done
}
