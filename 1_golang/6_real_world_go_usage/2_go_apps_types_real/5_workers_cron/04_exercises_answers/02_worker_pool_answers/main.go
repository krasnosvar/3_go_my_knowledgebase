// Упражнение 02: пул воркеров с сохранением порядка (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func process(inputs []int, workers int, fn func(int) int) []int {
	out := make([]int, len(inputs))
	idx := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				out[i] = fn(inputs[i])
			}
		}()
	}
	for i := range inputs {
		idx <- i
	}
	close(idx)
	wg.Wait()
	return out
}

func main() {
	fmt.Println("--- Проверка 02_worker_pool (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var active, peak atomic.Int32
	square := func(x int) int {
		n := active.Add(1)
		for {
			p := peak.Load()
			if n <= p || peak.CompareAndSwap(p, n) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		active.Add(-1)
		return x * x
	}

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := process(in, 3, square)
	check(fmt.Sprint(got) == "[1 4 9 16 25 36 49 64 81 100]", "результаты по порядку: "+fmt.Sprint(got))
	check(peak.Load() <= 3, fmt.Sprintf("не больше 3 воркеров одновременно (пик %d)", peak.Load()))
	check(peak.Load() >= 2, "работа действительно шла параллельно")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
