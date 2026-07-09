package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ready int32
	go func() { time.Sleep(5 * time.Millisecond); atomic.StoreInt32(&ready, 1) }()
	for atomic.LoadInt32(&ready) == 0 {
	}
	fmt.Println("ready")
}
