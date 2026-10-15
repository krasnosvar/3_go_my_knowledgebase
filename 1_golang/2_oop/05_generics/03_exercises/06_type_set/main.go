// Упражнение 06: type set и тильда (~).
//
// Задача: опишите constraint Integer как type set всех знаковых целых
// (с ~, чтобы подходили пользовательские типы вроде Level) и реализуйте:
//   - Abs(v) — модуль числа;
//   - Clamp(v, lo, hi) — ограничить v отрезком [lo, hi].
//
// Сейчас Integer содержит только int — тест с int8 и Level не скомпилируется,
// поэтому в заготовке он закомментирован: раскомментируйте после правки.
//
// Запуск: go run ./03_exercises/06_type_set/main.go
package main

import (
	"fmt"
)

// Level — пользовательский тип на базе int8.
type Level int8

// Integer — знаковые целые.
// TODO: ~int | ~int8 | ~int16 | ~int32 | ~int64
type Integer interface {
	~int
}

// Abs возвращает модуль.
func Abs[T Integer](v T) T {
	// TODO: если v < 0 — верните -v.
	return v
}

// Clamp ограничивает v отрезком [lo, hi].
func Clamp[T Integer](v, lo, hi T) T {
	// TODO
	return v
}

// extraChecks — проверки для типов, которые появятся в type set после правки.
// TODO: после расширения Integer раскомментируйте тело функции.
func extraChecks(check func(bool, string)) {
	check(false, "раскомментируйте extraChecks после расширения Integer")
	// check(Abs(int8(-5)) == 5, "Abs(int8(-5)): ожидалось 5")
	// check(Clamp(Level(9), 0, 5) == 5, "Clamp(Level(9), 0, 5): ожидалось 5 (работает благодаря ~int8)")
	// check(Abs(int64(-1<<40)) == 1<<40, "Abs(int64): ожидалось 1<<40")
}

func main() {
	fmt.Println("--- Проверка 06_type_set ---")

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
