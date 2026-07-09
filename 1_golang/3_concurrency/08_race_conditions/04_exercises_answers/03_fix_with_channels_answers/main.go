package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- 1
	}
	close(ch)
	s := 0
	for v := range ch {
		s += v
	}
	fmt.Println(s)
}
