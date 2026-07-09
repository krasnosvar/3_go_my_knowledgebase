package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	c := 0
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); mu.Lock(); c++; mu.Unlock() }()
	}
	wg.Wait()
	fmt.Println(c)
}
