package main

import (
	"fmt"
	"sync"
)

func main() {
	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(v int) { defer wg.Done(); sem <- struct{}{}; fmt.Println(v); <-sem }(i)
	}
	wg.Wait()
}
