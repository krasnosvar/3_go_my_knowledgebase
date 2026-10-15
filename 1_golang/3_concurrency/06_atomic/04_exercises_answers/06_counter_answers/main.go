// Упражнение 06: статистика запросов на атомиках (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type RequestStats struct {
	total  atomic.Int64
	errors atomic.Int64
}

func (s *RequestStats) Record(ok bool) {
	s.total.Add(1)
	if !ok {
		s.errors.Add(1)
	}
}

func (s *RequestStats) Snapshot() (total, errs int64) {
	return s.total.Load(), s.errors.Load()
}

func main() {
	fmt.Println("--- Проверка 06_counter (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var st RequestStats
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			st.Record(i%4 != 0) // каждый 4-й — ошибка
		}(i)
	}
	wg.Wait()
	total, errs := st.Snapshot()
	check(total == 100, fmt.Sprintf("total = %d, ожидалось 100", total))
	check(errs == 25, fmt.Sprintf("errors = %d, ожидалось 25", errs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
