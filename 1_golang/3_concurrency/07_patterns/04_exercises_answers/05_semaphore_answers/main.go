// Упражнение 05: ограничение параллелизма семафором (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func runLimited(tasks []func(), limit int) {
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			t()
		}(t)
	}
	wg.Wait()
}

// --- вспомогательные функции проверки (не меняйте) ---

// tracker считает текущее и максимальное число одновременных вызовов.
type tracker struct {
	active, max atomic.Int64
}

func (t *tracker) enter() {
	n := t.active.Add(1)
	for {
		m := t.max.Load()
		if n <= m || t.max.CompareAndSwap(m, n) {
			return
		}
	}
}

func (t *tracker) leave() { t.active.Add(-1) }

func main() {
	fmt.Println("--- Проверка 05_semaphore (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var tr tracker
	var done atomic.Int64
	tasks := make([]func(), 10)
	for i := range tasks {
		tasks[i] = func() {
			tr.enter()
			defer tr.leave()
			time.Sleep(5 * time.Millisecond)
			done.Add(1)
		}
	}
	runLimited(tasks, 3)
	check(done.Load() == 10, fmt.Sprintf("выполнено %d задач из 10", done.Load()))
	check(tr.max.Load() <= 3, fmt.Sprintf("одновременно выполнялось %d > 3 задач", tr.max.Load()))
	check(tr.max.Load() >= 2, "задачи выполнялись последовательно — семафор слишком строгий")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
