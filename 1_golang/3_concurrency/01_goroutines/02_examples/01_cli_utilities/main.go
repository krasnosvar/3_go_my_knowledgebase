package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, task := range []string{"fmt", "vet", "test"} {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			fmt.Println("run:", name)
		}(task)
	}
	wg.Wait()
}
