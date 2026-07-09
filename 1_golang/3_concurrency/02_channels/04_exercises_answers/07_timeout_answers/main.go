package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}
