// Упражнение 06: арифметика с time.Duration (Решение).
package main

import (
	"fmt"
	"time"
)

// sumDurations складывает длительности, заданные строками.
func sumDurations(ds []string) (time.Duration, error) {
	var total time.Duration
	for _, s := range ds {
		d, err := time.ParseDuration(s)
		if err != nil {
			return 0, fmt.Errorf("parse %q: %w", s, err)
		}
		total += d
	}
	return total, nil
}

func main() {
	fmt.Println("--- Проверка 06_duration_math (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	d, err := sumDurations([]string{"1h", "30m", "15s"})
	check(err == nil && d == time.Hour+30*time.Minute+15*time.Second, fmt.Sprintf("получено %v, %v", d, err))
	check(d.String() == "1h30m15s", fmt.Sprintf("String() = %q, ожидалось \"1h30m15s\"", d.String()))

	d, err = sumDurations([]string{"1m30s", "250ms", "250ms"})
	check(err == nil && d == 90*time.Second+500*time.Millisecond, fmt.Sprintf("получено %v, %v", d, err))

	_, err = sumDurations([]string{"10m", "5 minutes"})
	check(err != nil, "строка \"5 minutes\" должна дать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
