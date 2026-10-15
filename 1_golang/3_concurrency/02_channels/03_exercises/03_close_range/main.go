// Упражнение 03: закрытие канала и range.
//
// Задача:
//   - produce(n) возвращает канал, в который горутина отправляет 1..n и
//     затем ЗАКРЫВАЕТ его (закрывает всегда отправитель);
//   - sumAll(ch) читает канал через for v := range ch до закрытия и
//     возвращает сумму.
//
// Запуск: go run ./03_exercises/03_close_range/main.go
package main

import (
	"fmt"
	"time"
)

// produce отправляет 1..n в канал и закрывает его.
func produce(n int) <-chan int {
	// TODO: создайте канал, в горутине отправьте 1..n и закройте канал.
	return nil
}

// sumAll суммирует значения канала до его закрытия.
func sumAll(ch <-chan int) int {
	// TODO: for v := range ch { ... }
	return 0
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если канал не закрылся за d.
func collect(ch <-chan int, d time.Duration) ([]int, bool) {
	res := []int{}
	deadline := time.After(d)
	for {
		select {
		case v, open := <-ch:
			if !open {
				return res, true
			}
			res = append(res, v)
		case <-deadline:
			return res, false
		}
	}
}

func main() {
	fmt.Println("--- Проверка 03_close_range ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got, closed := collect(produce(4), 300*time.Millisecond)
	check(closed, "produce не закрыла канал")
	check(fmt.Sprint(got) == "[1 2 3 4]", fmt.Sprintf("produce(4) отдала %v, ожидалось [1 2 3 4]", got))

	ready := make(chan int, 3)
	ready <- 5
	ready <- 6
	ready <- 7
	close(ready)
	check(sumAll(ready) == 18, "sumAll([5 6 7]) должна вернуть 18")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
