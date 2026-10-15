// Упражнение 01: найти и устранить data race (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func safeCount(n int) int {
	var count atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
		}()
	}
	wg.Wait()
	return int(count.Load())
}

func main() {
	fmt.Println("--- Проверка 01_detect_race (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, n := range []int{1, 100, 1000} {
		got := safeCount(n)
		check(got == n, fmt.Sprintf("safeCount(%d) = %d", n, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
