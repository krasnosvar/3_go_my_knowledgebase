// Упражнение 08: consumer lag и алерт (Решение).
package main

import (
	"fmt"
	"sort"
)

func lagAlerts(end, committed map[int]int64, threshold int64) []string {
	type pl struct {
		p   int
		lag int64
	}
	var hot []pl
	for p, e := range end {
		if lag := e - committed[p]; lag > threshold {
			hot = append(hot, pl{p, lag})
		}
	}
	sort.Slice(hot, func(i, j int) bool {
		if hot[i].lag != hot[j].lag {
			return hot[i].lag > hot[j].lag
		}
		return hot[i].p < hot[j].p
	})
	out := make([]string, 0, len(hot))
	for _, h := range hot {
		out = append(out, fmt.Sprintf("p%d lag=%d", h.p, h.lag))
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 08_lag_alert (Решение) ---")

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
