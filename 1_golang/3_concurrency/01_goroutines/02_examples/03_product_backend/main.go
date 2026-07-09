package main

import (
	"fmt"
	"sync"
)

func fetch(id int) string { return fmt.Sprintf("item-%d", id) }

func main() {
	ids := []int{1, 2, 3}
	out := make([]string, len(ids))
	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			out[i] = fetch(id)
		}(i, id)
	}
	wg.Wait()
	fmt.Println(out)
}
