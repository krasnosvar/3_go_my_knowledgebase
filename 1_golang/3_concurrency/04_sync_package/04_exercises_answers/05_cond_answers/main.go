package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	ready := false
	go func() {
		mu.Lock()
		ready = true
		mu.Unlock()
		cond.Signal()
	}()
	mu.Lock()
	for !ready {
		cond.Wait()
	}
	mu.Unlock()
	fmt.Println("ready")
}
