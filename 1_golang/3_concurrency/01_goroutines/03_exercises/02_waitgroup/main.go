// Упражнение 02: ожидание горутин через sync.WaitGroup.
//
// Задача: реализуйте runAll(tasks) так, чтобы все задачи выполнялись
// конкурентно (каждая в своей горутине), а runAll возвращалась только
// после завершения ВСЕХ задач. Используйте sync.WaitGroup:
// wg.Add перед запуском, defer wg.Done() внутри горутины, wg.Wait() в конце.
//
// Проверка: 5 задач по 50ms должны завершиться заметно быстрее 250ms,
// и к моменту возврата runAll все 5 должны быть выполнены.
//
// Запуск: go run ./03_exercises/02_waitgroup/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// runAll выполняет все задачи конкурентно и ждёт их завершения.
func runAll(tasks []func()) {
	// TODO: запустите каждую задачу в горутине и дождитесь всех через WaitGroup.
}

func main() {
	fmt.Println("--- Проверка 02_waitgroup ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var finished atomic.Int64
	tasks := make([]func(), 5)
	for i := range tasks {
		tasks[i] = func() {
			time.Sleep(50 * time.Millisecond)
			finished.Add(1)
		}
	}

	start := time.Now()
	runAll(tasks)
	elapsed := time.Since(start)

	check(finished.Load() == 5, fmt.Sprintf("после runAll завершено %d задач из 5", finished.Load()))
	check(elapsed < 200*time.Millisecond, fmt.Sprintf("задачи выполнялись последовательно (%v), нужен параллельный запуск", elapsed))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
