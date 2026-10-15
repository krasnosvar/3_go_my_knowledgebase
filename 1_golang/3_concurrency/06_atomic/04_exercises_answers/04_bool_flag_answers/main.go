// Упражнение 04: флаг закрытия на atomic.Bool (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Shutdown struct {
	closed atomic.Bool
}

func (s *Shutdown) Close() bool {
	return s.closed.CompareAndSwap(false, true)
}

func (s *Shutdown) IsClosed() bool {
	return s.closed.Load()
}

func main() {
	fmt.Println("--- Проверка 04_bool_flag (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Shutdown
	check(!s.IsClosed(), "новый флаг не должен быть закрыт")

	var winners atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if s.Close() {
				winners.Add(1)
			}
		}()
	}
	wg.Wait()
	check(winners.Load() == 1, fmt.Sprintf("Close вернула true %d раз, ожидалось ровно 1", winners.Load()))
	check(s.IsClosed(), "после Close флаг должен быть закрыт")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
