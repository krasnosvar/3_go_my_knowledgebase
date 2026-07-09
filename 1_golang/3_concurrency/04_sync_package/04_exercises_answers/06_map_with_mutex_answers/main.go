package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			mu.Lock()
			m[v] = v * v
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
