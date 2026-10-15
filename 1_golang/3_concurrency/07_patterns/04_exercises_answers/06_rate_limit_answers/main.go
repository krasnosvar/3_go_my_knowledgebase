// Упражнение 06: ограничение частоты через time.Ticker (Решение).
package main

import (
	"fmt"
	"time"
)

func throttle(n int, interval time.Duration, do func(i int)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for i := 0; i < n; i++ {
		<-ticker.C
		do(i)
	}
}

func main() {
	fmt.Println("--- Проверка 06_rate_limit (Решение) ---")

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
