// Упражнение 01: атомарный инкремент (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func countConcurrently(n int) int64 {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&counter)
}

func main() {
	fmt.Println("--- Проверка 01_add (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countConcurrently(50)
	check(got == 5000, fmt.Sprintf("countConcurrently(50) = %d, ожидалось 5000", got))
	check(countConcurrently(0) == 0, "для 0 горутин ожидался 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
