package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var reads int64
	m := map[string]int{"a": 1}
	var mu sync.Mutex
	atomic.AddInt64(&reads, 1)
	mu.Lock()
	m["b"] = 2
	mu.Unlock()
	fmt.Println(reads, len(m))
}
