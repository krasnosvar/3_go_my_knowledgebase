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

// Append добавляет число в слайс, защищая запись с помощью мьютекса
func (sc *SafeCollector) Append(val int) {
	// TODO: Реализуйте защиту записи с помощью sc.mu. Lock / Unlock
}

func main() {
	fmt.Println("--- Проверка 08_concurrency_safety ---")

	collector := &SafeCollector{
		data: make([]int, 0),
	}

	var wg sync.WaitGroup
	workers := 100
	wg.Add(workers)

	// Запускаем 100 конкурентных воркеров, каждый пишет по одному числу
	for i := 0; i < workers; i++ {
		go func(val int) {
			defer wg.Done()
			collector.Append(val)
		}(i)
	}

	wg.Wait()

	// Если мьютекс реализован верно, длина слайса будет ровно 100.
	// Без мьютекса длина будет меньше из-за race condition или произойдет краш.
	if len(collector.data) == 100 {
		fmt.Println("Тест конкурентного слайса: PASS")
	} else {
		fmt.Printf("Тест конкурентного слайса: FAIL (записано %d из %d)\n", len(collector.data), workers)
	}
}
