// Упражнение 03: CAS-цикл: атомарный максимум.
//
// Задача: casMax(addr, v) атомарно обновляет *addr до max(*addr, v).
// Шаблон CAS-цикла:
//
//	for {
//	    old := atomic.LoadInt64(addr)
//	    if v <= old { return }
//	    if atomic.CompareAndSwapInt64(addr, old, v) { return }
//	    // кто-то успел изменить значение — пробуем снова
//	}
//
// Запуск: go run ./03_exercises/03_compare_swap/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// casMax атомарно записывает v, если оно больше текущего значения.
func casMax(addr *int64, v int64) {
	// TODO: CAS-цикл.
}

func main() {
	fmt.Println("--- Проверка 03_compare_swap ---")

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
