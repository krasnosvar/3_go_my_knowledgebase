// Упражнение 07: временная настройка GOMAXPROCS (Решение).
package main

import (
	"fmt"
	"runtime"
)

// withGOMAXPROCS выполняет f при GOMAXPROCS == n и восстанавливает старое значение.
func withGOMAXPROCS(n int, f func()) {
	old := runtime.GOMAXPROCS(n)
	defer runtime.GOMAXPROCS(old)
	f()
}

func main() {
	fmt.Println("--- Проверка 07_gomaxprocs_tuning (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	orig := runtime.GOMAXPROCS(0)
	target := 1
	if orig == 1 {
		target = 2
	}

	var inside int
	withGOMAXPROCS(target, func() { inside = runtime.GOMAXPROCS(0) })
	check(inside == target, fmt.Sprintf("внутри f GOMAXPROCS должен быть %d, получено %d", target, inside))
	check(runtime.GOMAXPROCS(0) == orig, fmt.Sprintf("после f GOMAXPROCS должен вернуться к %d, получено %d", orig, runtime.GOMAXPROCS(0)))

	func() {
		defer func() { _ = recover() }()
		withGOMAXPROCS(target, func() { panic("boom") })
	}()
	check(runtime.GOMAXPROCS(0) == orig, "после паники в f GOMAXPROCS тоже должен восстановиться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
