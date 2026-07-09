package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("no data")
	}
}
