package main

import (
	"context"
	"fmt"
)

func main() {
	parent, cancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	cancel()
	<-child.Done()
	fmt.Println(child.Err())
}
