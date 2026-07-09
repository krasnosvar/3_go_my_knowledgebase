package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() { close(done) }()
	<-done
	fmt.Println("no deadlock")
}
