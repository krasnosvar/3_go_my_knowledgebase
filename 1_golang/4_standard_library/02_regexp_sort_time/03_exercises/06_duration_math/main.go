// Упражнение 06: арифметика с time.Duration.
//
// Задача: реализуйте sumDurations(ds), которая разбирает строки вида
// "1h", "30m", "1m30s", "250ms" через time.ParseDuration и возвращает сумму.
//   - любая неверная строка → (0, ошибка с указанием этой строки, %w);
//   - пустой список → 0, nil.
//
// Запуск: go run ./03_exercises/06_duration_math/main.go
package main

import (
	"fmt"
	"time"
)

// sumDurations складывает длительности, заданные строками.
func sumDurations(ds []string) (time.Duration, error) {
	// TODO: time.ParseDuration для каждой строки, накапливайте сумму.
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 06_duration_math ---")

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
