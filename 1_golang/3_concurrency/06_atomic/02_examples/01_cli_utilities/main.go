package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var calls int64
	atomic.AddInt64(&calls, 1)
	fmt.Println(calls)
}
