// Упражнение 05: fan-out: несколько worker'ов читают один канал (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func fanOut(jobs []int, workers int, f func(int) int) []int {
	results := make([]int, len(jobs))
	idx := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				results[i] = f(jobs[i])
			}
		}()
	}
	for i := range jobs {
		idx <- i
	}
	close(idx)
	wg.Wait()
	return results
}

func main() {
	fmt.Println("--- Проверка 05_fan_out (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var active, maxActive atomic.Int64
	f := func(x int) int {
		n := active.Add(1)
		for {
			m := maxActive.Load()
			if n <= m || maxActive.CompareAndSwap(m, n) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		active.Add(-1)
		return x * 10
	}

	got := fanOut([]int{1, 2, 3, 4, 5, 6}, 3, f)
	check(fmt.Sprint(got) == "[10 20 30 40 50 60]", fmt.Sprintf("fanOut = %v, ожидалось [10 20 30 40 50 60]", got))
	check(maxActive.Load() >= 2, "задачи выполнялись последовательно, а не несколькими worker'ами")
	check(maxActive.Load() <= 3, fmt.Sprintf("одновременно работало %d > 3 worker'ов", maxActive.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
