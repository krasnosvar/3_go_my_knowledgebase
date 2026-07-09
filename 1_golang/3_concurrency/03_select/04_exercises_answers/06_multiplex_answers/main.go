package main

import "fmt"

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)
	a <- 10
	b <- 20
	close(a)
	close(b)
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				break
			}
			fmt.Println(v)
		case v, ok := <-b:
			if !ok {
				b = nil
				break
			}
			fmt.Println(v)
		}
	}
}
