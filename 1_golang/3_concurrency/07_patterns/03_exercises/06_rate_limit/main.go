// Упражнение 06: ограничение частоты через time.Ticker.
//
// Задача: throttle(n, interval, do) вызывает do(i) для i = 0..n-1 не чаще
// одного раза в interval. Используйте time.NewTicker(interval) и перед
// каждым вызовом ждите <-ticker.C. Не забудьте defer ticker.Stop().
//
// Это простейший rate limiter; для всплесков (burst) применяют
// token bucket — см. golang.org/x/time/rate.
//
// Запуск: go run ./03_exercises/06_rate_limit/main.go
package main

import (
	"fmt"
	"time"
)

// throttle вызывает do n раз с интервалом не меньше interval.
func throttle(n int, interval time.Duration, do func(i int)) {
	// TODO: ticker := time.NewTicker(interval); defer ticker.Stop(); ...
}

func main() {
	fmt.Println("--- Проверка 06_rate_limit ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var stamps []time.Time
	var ids []int
	throttle(5, 10*time.Millisecond, func(i int) {
		stamps = append(stamps, time.Now())
		ids = append(ids, i)
	})
	check(fmt.Sprint(ids) == "[0 1 2 3 4]", fmt.Sprintf("вызовы = %v, ожидалось [0 1 2 3 4]", ids))
	for i := 1; i < len(stamps); i++ {
		gap := stamps[i].Sub(stamps[i-1])
		check(gap >= 7*time.Millisecond, fmt.Sprintf("между вызовами %d и %d прошло %v < interval", i-1, i, gap))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
