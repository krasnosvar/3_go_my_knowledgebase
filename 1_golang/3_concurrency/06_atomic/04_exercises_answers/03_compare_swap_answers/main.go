// Упражнение 03: CAS-цикл: атомарный максимум (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func casMax(addr *int64, v int64) {
	for {
		old := atomic.LoadInt64(addr)
		if v <= old {
			return
		}
		if atomic.CompareAndSwapInt64(addr, old, v) {
			return
		}
	}
}

func main() {
	fmt.Println("--- Проверка 03_compare_swap (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var maxVal int64
	var wg sync.WaitGroup
	for i := int64(0); i < 100; i++ {
		wg.Add(1)
		go func(v int64) {
			defer wg.Done()
			casMax(&maxVal, v)
		}(i)
	}
	wg.Wait()
	check(atomic.LoadInt64(&maxVal) == 99, fmt.Sprintf("max = %d, ожидалось 99", atomic.LoadInt64(&maxVal)))

	casMax(&maxVal, 5)
	check(atomic.LoadInt64(&maxVal) == 99, "меньшее значение не должно перезаписывать максимум")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
