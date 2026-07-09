package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{"a": 1}
	var rw sync.RWMutex
	rw.RLock()
	fmt.Println(m["a"])
	rw.RUnlock()
	rw.Lock()
	m["b"] = 2
	rw.Unlock()
	rw.RLock()
	fmt.Println(m["b"])
	rw.RUnlock()
}
