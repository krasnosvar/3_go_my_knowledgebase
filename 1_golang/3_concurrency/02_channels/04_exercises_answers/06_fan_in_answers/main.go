package main

import "fmt"

func fanIn(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					break
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					break
				}
				out <- v
			}
		}
	}()
	return out
}

func main() {
	a := make(chan int, 2)
	b := make(chan int, 2)
	a <- 1
	a <- 2
	close(a)
	b <- 10
	b <- 20
	close(b)
	for v := range fanIn(a, b) {
		fmt.Println(v)
	}
}
