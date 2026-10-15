// Упражнение 01: worker pool фиксированного размера.
//
// Задача: runPool(jobs, workers, f) обрабатывает jobs пулом из ровно
// workers горутин: канал задач (индексы), workers горутин читают его
// через range, пишут results[i] = f(jobs[i]); после отправки всех задач
// канал закрывается, функция ждёт пул (WaitGroup) и возвращает results.
//
// Тест проверяет порядок результатов и что одновременно работало не
// больше workers горутин (и больше одной).
//
// Запуск: go run ./03_exercises/01_worker_pool/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// runPool обрабатывает jobs пулом из workers горутин.
func runPool(jobs []int, workers int, f func(int) int) []int {
	// TODO: канал индексов, workers горутин, close, wg.Wait().
	return nil
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
	fmt.Println("--- Проверка 01_worker_pool ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var tr tracker
	f := func(x int) int {
		tr.enter()
		defer tr.leave()
		time.Sleep(3 * time.Millisecond)
		return x + 100
	}
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := runPool(jobs, 3, f)
	check(fmt.Sprint(got) == "[101 102 103 104 105 106 107 108 109 110]", fmt.Sprintf("results = %v", got))
	check(tr.max.Load() <= 3, fmt.Sprintf("одновременно работало %d > 3 горутин", tr.max.Load()))
	check(tr.max.Load() >= 2, "задачи выполнялись последовательно — пул не работает")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
