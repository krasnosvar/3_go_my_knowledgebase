// Упражнение 07: цикл на time.Ticker с остановкой (Решение).
package main

import (
	"fmt"
	"time"
)

// collectTicks ждёт n тиков и вызывает onTick для каждого.
func collectTicks(n int, interval time.Duration, onTick func(i int)) int {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	count := 0
	for count < n {
		<-ticker.C
		count++
		onTick(count)
	}
	return count
}

func main() {
	fmt.Println("--- Проверка 07_ticker_loop (Решение) ---")

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
