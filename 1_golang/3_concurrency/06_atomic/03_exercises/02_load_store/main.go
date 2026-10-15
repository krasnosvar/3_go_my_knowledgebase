// Упражнение 02: атомарные Load и Store.
//
// Задача: Status хранит текущее состояние сервиса (int32), которое одни
// горутины пишут, а другие читают. Реализуйте Set через v.Store и Get
// через v.Load на типе atomic.Int32 (Go 1.19+). Ни чтение, ни запись
// обычного поля без синхронизации здесь недопустимы — это гонка.
//
// Запуск: go run ./03_exercises/02_load_store/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Status — атомарно читаемое/записываемое состояние.
type Status struct {
	v atomic.Int32
}

// Set атомарно записывает состояние.
func (s *Status) Set(code int32) {
	// TODO: s.v.Store(code)
}

// Get атомарно читает состояние.
func (s *Status) Get() int32 {
	// TODO: return s.v.Load()
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_load_store ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Status
	var wg sync.WaitGroup
	var bad atomic.Int64
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			s.Set(int32(i % 3))
		}(i)
		go func() {
			defer wg.Done()
			if v := s.Get(); v < 0 || v > 2 {
				bad.Add(1)
			}
		}()
	}
	wg.Wait()
	check(bad.Load() == 0, "прочитано значение вне диапазона 0..2")

	s.Set(7)
	check(s.Get() == 7, fmt.Sprintf("после Set(7) Get = %d", s.Get()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
