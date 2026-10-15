// Упражнение 01: worker pool фиксированного размера (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func runPool(jobs []int, workers int, f func(int) int) []int {
	results := make([]int, len(jobs))
	tasks := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range tasks {
				results[i] = f(jobs[i])
			}
		}()
	}
	for i := range jobs {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	return results
}

// --- вспомогательные функции проверки (не меняйте) ---

// tracker считает текущее и максимальное число одновременных вызовов.
type tracker struct {
	active, max atomic.Int64
}

func (t *tracker) enter() {
	n := t.active.Add(1)
	for {
		m := t.max.Load()
		if n <= m || t.max.CompareAndSwap(m, n) {
			return
		}
	}
}

func (t *tracker) leave() { t.active.Add(-1) }

func main() {
	fmt.Println("--- Проверка 01_worker_pool (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var tr tracker
	f := func(x int) int {
		tr.enter()
		defer tr.leave()
		time.Sleep(3 * time.Millisecond)
		return x + 100
	}
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := runPool(jobs, 3, f)
	check(fmt.Sprint(got) == "[101 102 103 104 105 106 107 108 109 110]", fmt.Sprintf("results = %v", got))
	check(tr.max.Load() <= 3, fmt.Sprintf("одновременно работало %d > 3 горутин", tr.max.Load()))
	check(tr.max.Load() >= 2, "задачи выполнялись последовательно — пул не работает")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
