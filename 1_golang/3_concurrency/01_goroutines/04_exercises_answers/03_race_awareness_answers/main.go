// Упражнение 03: общий счётчик без гонки данных (Решение).
package main

import (
	"fmt"
	"sync"
)

func countParallel(workers, perWorker int) int {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		counter int
	)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perWorker; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}

func main() {
	fmt.Println("--- Проверка 03_race_awareness (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countParallel(50, 1000)
	check(got == 50000, fmt.Sprintf("countParallel(50, 1000) = %d, ожидалось 50000", got))

	got = countParallel(1, 10)
	check(got == 10, fmt.Sprintf("countParallel(1, 10) = %d, ожидалось 10", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
