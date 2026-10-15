// Упражнение 04: наблюдение за циклами GC (Решение).
package main

import (
	"fmt"
	"runtime"
)

// gcCyclesDuring возвращает число завершённых циклов GC за время f().
func gcCyclesDuring(f func()) uint32 {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before) // ReadMemStats делает короткую STW-паузу — не для hot path
	f()
	runtime.ReadMemStats(&after)
	return after.NumGC - before.NumGC
}

func main() {
	fmt.Println("--- Проверка 04_gc_observation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	two := gcCyclesDuring(func() {
		runtime.GC()
		runtime.GC()
	})
	check(two >= 2, fmt.Sprintf("два runtime.GC() — ожидалось >= 2 циклов, получено %d", two))

	one := gcCyclesDuring(runtime.GC)
	check(one >= 1, fmt.Sprintf("один runtime.GC() — ожидалось >= 1 цикла, получено %d", one))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
