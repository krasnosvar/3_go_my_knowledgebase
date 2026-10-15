// ⏩ Забегание вперёд: горутины и sync.Mutex разбираются в 3_concurrency (01_goroutines, 04_sync_package).

package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	stats map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		stats: make(map[string]int),
	}
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stats[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.stats[key]
}

func main() {
	fmt.Println("--- Проверка 06_maps_concurrency (Решение) ---")

	counter := NewSafeCounter()

	var wg sync.WaitGroup
	workers := 100
	wg.Add(workers)

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
