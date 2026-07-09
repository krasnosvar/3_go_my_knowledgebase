package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 7
	select {
	case v := <-ch:
		fmt.Println(v)
	}
}
