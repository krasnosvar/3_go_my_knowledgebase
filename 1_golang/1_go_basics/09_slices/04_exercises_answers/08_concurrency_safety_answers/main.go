// ⏩ Забегание вперёд: горутины и sync.Mutex разбираются в 3_concurrency (01_goroutines, 04_sync_package);
// методы — в 2_oop/01_methods.

package main

import (
	"fmt"
	"sync"
)

// SafeCollector — потокобезопасный накопитель чисел
type SafeCollector struct {
	mu   sync.Mutex
	data []int
}

// Append добавляет число в слайс, используя Mutex Lock/Unlock
func (sc *SafeCollector) Append(val int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.data = append(sc.data, val)
}

func main() {
	fmt.Println("--- Проверка 08_concurrency_safety (Решение) ---")

	collector := &SafeCollector{
		data: make([]int, 0),
	}

	var wg sync.WaitGroup
	workers := 100
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(val int) {
			defer wg.Done()
			collector.Append(val)
		}(i)
	}

	wg.Wait()

	if len(collector.data) == 100 {
		fmt.Println("Тест конкурентного слайса: PASS")
	} else {
		fmt.Printf("Тест конкурентного слайса: FAIL (записано %d из %d)\n", len(collector.data), workers)
	}
}
