package main

import (
	"fmt"
	"sync"
)

func main() {
	status := map[string]bool{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, node := range []string{"n1", "n2", "n3"} {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			mu.Lock()
			status[n] = true
			mu.Unlock()
		}(node)
	}
	wg.Wait()
	fmt.Println(status)
}
