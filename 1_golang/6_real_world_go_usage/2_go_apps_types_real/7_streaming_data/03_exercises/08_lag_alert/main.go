// Упражнение 08: consumer lag и алерт.
//
// Lag партиции = endOffset - committed: сколько сообщений ещё не обработано.
//
// Задача: реализуйте lagAlerts(end, committed map[int]int64, threshold int64) []string:
//   - для каждой партиции из end посчитать lag (нет коммита → committed = 0);
//   - вернуть строки "p<N> lag=<L>" для партиций с lag > threshold,
//     отсортированные по убыванию lag, при равенстве — по номеру партиции.
//
// Запуск: go run ./03_exercises/08_lag_alert
package main

import (
	"fmt"
	"sort"
)

var _ = sort.Slice

func lagAlerts(end, committed map[int]int64, threshold int64) []string {
	// TODO: посчитайте lag, отфильтруйте, отсортируйте, отформатируйте.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_lag_alert ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	end := map[int]int64{0: 1000, 1: 500, 2: 800, 3: 300, 4: 700}
	committed := map[int]int64{0: 990, 1: 100, 2: 400, 4: 300}
	got := fmt.Sprint(lagAlerts(end, committed, 100))
	check(got == "[p1 lag=400 p2 lag=400 p4 lag=400 p3 lag=300]", "алерты по убыванию lag: "+got)
	check(len(lagAlerts(end, committed, 1000)) == 0, "высокий порог → алертов нет")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
