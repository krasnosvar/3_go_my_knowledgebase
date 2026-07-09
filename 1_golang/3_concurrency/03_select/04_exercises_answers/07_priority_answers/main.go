package main

import "fmt"

func main() {
	urgent := make(chan string, 1)
	normal := make(chan string, 1)
	urgent <- "u1"
	normal <- "n1"
	select {
	case v := <-urgent:
		fmt.Println("urgent", v)
	default:
		select {
		case v := <-normal:
			fmt.Println("normal", v)
		}
	}
}
