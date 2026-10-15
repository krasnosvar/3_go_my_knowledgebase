// Упражнение 06: type set и тильда (~) (Решение).
package main

import (
	"fmt"
)

// Level — пользовательский тип на базе int8.
type Level int8

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Abs[T Integer](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func Clamp[T Integer](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func extraChecks(check func(bool, string)) {
	check(Abs(int8(-5)) == 5, "Abs(int8(-5)): ожидалось 5")
	check(Clamp(Level(9), 0, 5) == 5, "Clamp(Level(9), 0, 5): ожидалось 5 (работает благодаря ~int8)")
	check(Abs(int64(-1<<40)) == 1<<40, "Abs(int64): ожидалось 1<<40")
}

func main() {
	fmt.Println("--- Проверка 06_type_set (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(Abs(-7) == 7, fmt.Sprintf("Abs(-7): ожидалось 7, получено %d", Abs(-7)))
	check(Abs(3) == 3, "Abs(3): ожидалось 3")
	check(Clamp(15, 0, 10) == 10, fmt.Sprintf("Clamp(15,0,10): ожидалось 10, получено %d", Clamp(15, 0, 10)))
	check(Clamp(-3, 0, 10) == 0, fmt.Sprintf("Clamp(-3,0,10): ожидалось 0, получено %d", Clamp(-3, 0, 10)))
	check(Clamp(5, 0, 10) == 5, "Clamp(5,0,10): ожидалось 5")
	extraChecks(check)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
