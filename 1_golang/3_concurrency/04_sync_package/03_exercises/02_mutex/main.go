// Упражнение 02: счётчик под sync.Mutex.
//
// Задача: сделайте Counter безопасным для конкурентного использования:
// Inc() увеличивает значение на 1, Value() возвращает текущее. Оба метода
// захватывают c.mu (Lock + defer Unlock). Ресивер — указатель: мьютекс
// нельзя копировать.
//
// Проверьте себя: go run -race ./03_exercises/02_mutex/main.go
//
// Запуск: go run ./03_exercises/02_mutex/main.go
package main

import (
	"fmt"
	"sync"
)

// Counter — конкурентно-безопасный счётчик.
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc увеличивает счётчик на 1.
func (c *Counter) Inc() {
	// TODO: c.mu.Lock(); defer c.mu.Unlock(); c.n++
}

// Value возвращает текущее значение.
func (c *Counter) Value() int {
	// TODO: прочитайте c.n под мьютексом.
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_mutex ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	check(c.Value() == 10000, fmt.Sprintf("Value = %d, ожидалось 10000", c.Value()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
