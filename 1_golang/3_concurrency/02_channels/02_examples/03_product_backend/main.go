package main

import "fmt"

func producer(out chan<- int) {
	defer close(out)
	for i := 1; i <= 3; i++ {
		out <- i
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("value", v)
	}
}
