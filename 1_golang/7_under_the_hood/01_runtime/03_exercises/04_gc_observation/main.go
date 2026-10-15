// Упражнение 04: наблюдение за циклами GC.
//
// Задача: реализуйте gcCyclesDuring(f), которая возвращает, сколько циклов
// сборки мусора завершилось за время выполнения f. Используйте
// runtime.ReadMemStats и поле MemStats.NumGC (до и после вызова f).
//
// Проверка: f с двумя runtime.GC() даёт не меньше 2 циклов, с одним — не
// меньше 1. Точное число не проверяется: фоновый GC может добавить циклы.
//
// Запуск: go run ./03_exercises/04_gc_observation/main.go
package main

import (
	"fmt"
	"runtime"
)

// gcCyclesDuring возвращает число завершённых циклов GC за время f().
func gcCyclesDuring(f func()) uint32 {
	// TODO: прочитайте MemStats.NumGC до f(), вызовите f(), прочитайте снова
	// и верните разницу.
	f()
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_gc_observation ---")

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
