// Упражнение 05: ограничение параллелизма семафором.
//
// Задача: runLimited(tasks, limit) запускает каждую задачу в своей
// горутине, но одновременно выполняется не больше limit задач. Семафор —
// буферизованный канал ёмкостью limit:
//
//	sem <- struct{}{}        // занять слот (блокирует, если все заняты)
//	defer func() { <-sem }() // освободить
//
// Функция возвращается после завершения всех задач.
//
// Запуск: go run ./03_exercises/05_semaphore/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// runLimited выполняет tasks с параллелизмом не выше limit.
func runLimited(tasks []func(), limit int) {
	// TODO: sem := make(chan struct{}, limit); горутина на задачу; wg.Wait().
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
	fmt.Println("--- Проверка 05_semaphore ---")

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
