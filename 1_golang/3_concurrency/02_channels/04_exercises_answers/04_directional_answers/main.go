package main

import "fmt"

func sender(out chan<- int) {
	out <- 9
	close(out)
}

func receiver(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}
