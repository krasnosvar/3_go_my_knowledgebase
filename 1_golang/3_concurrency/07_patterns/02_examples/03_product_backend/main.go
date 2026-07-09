package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)
	go func() {
		defer close(in)
		for _, v := range []int{1, 2, 3} {
			in <- v
		}
	}()
	go func() {
		defer close(out)
		for v := range in {
			out <- v * 10
		}
	}()
	for v := range out {
		fmt.Println(v)
	}
}
