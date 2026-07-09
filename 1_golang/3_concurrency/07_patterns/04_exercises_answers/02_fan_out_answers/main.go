package main

import "fmt"

func main() {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)
	for j := range jobs {
		fmt.Println(j)
	}
}
