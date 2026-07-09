package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{}
	var mu sync.Mutex
	mu.Lock()
	m["a"] = 1
	mu.Unlock()
	fmt.Println(m["a"])
}
