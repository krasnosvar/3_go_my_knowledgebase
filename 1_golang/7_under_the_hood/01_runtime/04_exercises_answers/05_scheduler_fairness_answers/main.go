// Упражнение 05: справедливость планировщика (асинхронное вытеснение) (Решение).
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// worker крутит тугой цикл, увеличивая counters[id], пока stop == false.
func worker(id int, counters []int64, stop *atomic.Bool) {
	for !stop.Load() {
		atomic.AddInt64(&counters[id], 1)
	}
}

func main() {
	fmt.Println("--- Проверка 05_scheduler_fairness (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	old := runtime.GOMAXPROCS(1) // один P: горутины делят одно «ядро»
	defer runtime.GOMAXPROCS(old)

	const n = 4
	counters := make([]int64, n)
	var stop atomic.Bool
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, counters, &stop)
		}(i)
	}

	time.Sleep(100 * time.Millisecond) // main тоже должен получить P обратно
	stop.Store(true)
	wg.Wait()

	for i := 0; i < n; i++ {
		c := atomic.LoadInt64(&counters[i])
		check(c > 0, fmt.Sprintf("воркер %d не получил процессорного времени (counter=%d)", i, c))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
