// ⏩ Забегание вперёд: горутины и sync.Mutex разбираются в 3_concurrency (01_goroutines, 04_sync_package);
// методы — в 2_oop/01_methods. Эту задачу можно отложить до прохождения этих тем.

package main

import (
	"fmt"
	"sync"
)

// SafeCounter — потокобезопасный счетчик на основе map и sync.Mutex
type SafeCounter struct {
	mu    sync.Mutex
	stats map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		stats: make(map[string]int),
	}
}

// Inc увеличивает счетчик для указанного ключа на 1
func (c *SafeCounter) Inc(key string) {
	// TODO: Реализуйте потокобезопасное увеличение счетчика.
}

// Value возвращает текущее значение счетчика для ключа
func (c *SafeCounter) Value(key string) int {
	// TODO: Реализуйте потокобезопасное чтение.
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_maps_concurrency ---")

	counter := NewSafeCounter()

	var wg sync.WaitGroup
	workers := 100
	wg.Add(workers)

	// Конкурентно инкрементируем ключ "hits"
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			counter.Inc("hits")
		}()
	}

	wg.Wait()

	val := counter.Value("hits")
	if val == workers {
		fmt.Println("Тест SafeCounter: PASS")
	} else {
		fmt.Printf("Тест SafeCounter: FAIL (ожидалось %d, получено %d)\n", workers, val)
	}
}
