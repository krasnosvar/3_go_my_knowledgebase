package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(id int) {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("worker", id)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
