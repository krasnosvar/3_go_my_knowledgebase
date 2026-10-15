// Упражнение 02: ожидание горутин через sync.WaitGroup (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func runAll(tasks []func()) {
	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
		}(t)
	}
	wg.Wait()
}

func main() {
	fmt.Println("--- Проверка 02_waitgroup (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var finished atomic.Int64
	tasks := make([]func(), 5)
	for i := range tasks {
		tasks[i] = func() {
			time.Sleep(50 * time.Millisecond)
			finished.Add(1)
		}
	}

	start := time.Now()
	runAll(tasks)
	elapsed := time.Since(start)

	check(finished.Load() == 5, fmt.Sprintf("после runAll завершено %d задач из 5", finished.Load()))
	check(elapsed < 200*time.Millisecond, fmt.Sprintf("задачи выполнялись последовательно (%v), нужен параллельный запуск", elapsed))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
