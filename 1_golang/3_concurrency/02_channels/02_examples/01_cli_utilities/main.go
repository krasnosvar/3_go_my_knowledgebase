package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { ch <- "build ok" }()
	fmt.Println(<-ch)
}
