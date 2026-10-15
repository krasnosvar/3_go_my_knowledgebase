// Упражнение 07: цикл на time.Ticker с остановкой.
//
// Задача: реализуйте collectTicks(n, interval), которая:
//   - создаёт time.NewTicker(interval) и ОБЯЗАТЕЛЬНО останавливает его
//     (defer ticker.Stop());
//   - в цикле читает ticker.C ровно n раз, вызывая onTick(i) с номером тика 1..n;
//   - возвращает число полученных тиков.
//
// Запуск: go run ./03_exercises/07_ticker_loop/main.go
package main

import (
	"fmt"
	"time"
)

// collectTicks ждёт n тиков и вызывает onTick для каждого.
func collectTicks(n int, interval time.Duration, onTick func(i int)) int {
	// TODO: NewTicker + defer Stop + цикл из n чтений <-ticker.C.
	return 0
}

func main() {
	fmt.Println("--- Проверка 07_ticker_loop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var seen []int
	start := time.Now()
	got := collectTicks(3, 10*time.Millisecond, func(i int) { seen = append(seen, i) })
	elapsed := time.Since(start)

	check(got == 3, fmt.Sprintf("ожидалось 3 тика, получено %d", got))
	check(fmt.Sprint(seen) == "[1 2 3]", fmt.Sprintf("onTick вызван с %v, ожидалось [1 2 3]", seen))
	check(elapsed >= 30*time.Millisecond, fmt.Sprintf("3 тика по 10ms не могут пройти за %v", elapsed))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
