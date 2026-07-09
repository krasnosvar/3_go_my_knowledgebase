package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { fmt.Println("hello") }()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main done")
}
