package main

import "fmt"

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)
	out := make(chan int, 2)
	a <- 1
	b <- 2
	close(a)
	close(b)
	for v := range a {
		out <- v
	}
	for v := range b {
		out <- v
	}
	close(out)
	for v := range out {
		fmt.Println(v)
	}
}
