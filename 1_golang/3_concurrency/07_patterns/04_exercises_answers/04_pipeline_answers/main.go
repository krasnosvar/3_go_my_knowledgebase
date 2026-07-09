package main

import "fmt"

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for v := range in {
		fmt.Println(v * 2)
	}
}
