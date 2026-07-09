package main

import "fmt"

func mul2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}()
	return out
}

func onlyGt4(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v > 4 {
				out <- v
			}
		}
	}()
	return out
}

func main() {
	in := make(chan int)
	go func() {
		defer close(in)
		for _, v := range []int{1, 2, 3} {
			in <- v
		}
	}()
	for v := range onlyGt4(mul2(in)) {
		fmt.Println(v)
	}
}
