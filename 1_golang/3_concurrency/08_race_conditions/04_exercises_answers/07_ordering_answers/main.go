package main

import "fmt"

func main() {
	ready := make(chan struct{})
	go func() { close(ready) }()
	<-ready
	fmt.Println("ordered")
}
