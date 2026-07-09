package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, cmd := range []string{"fmt", "vet", "test"} {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			fmt.Println("run", c)
		}(cmd)
	}
	wg.Wait()
}
