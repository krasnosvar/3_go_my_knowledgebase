// Упражнение 08: барьер между фазами.
//
// Задача: runPhases(workers, phase1, phase2) запускает workers горутин.
// Каждая выполняет phase1(id), затем ждёт, пока ВСЕ горутины закончат
// phase1 (барьер), и только потом выполняет phase2(id). Функция
// возвращается после завершения phase2 у всех.
//
// Простой барьер: WaitGroup для фазы 1 + канал start, который закрывается
// после wg1.Wait(); вторая WaitGroup ждёт фазу 2.
//
// Запуск: go run ./03_exercises/08_barrier/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// runPhases выполняет phase1 у всех, затем phase2 у всех.
func runPhases(workers int, phase1, phase2 func(id int)) {
	// TODO: wg1 для фазы 1, канал-барьер, wg2 для фазы 2.
}

func main() {
	fmt.Println("--- Проверка 08_barrier ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	const workers = 5
	var p1, p2 atomic.Int64
	var early atomic.Int64
	runPhases(workers,
		func(id int) {
			time.Sleep(time.Duration(id) * time.Millisecond) // разная длительность фазы 1
			p1.Add(1)
		},
		func(id int) {
			if p1.Load() != workers {
				early.Add(1)
			}
			p2.Add(1)
		},
	)
	check(p1.Load() == workers, fmt.Sprintf("phase1 выполнена %d раз, ожидалось %d", p1.Load(), workers))
	check(p2.Load() == workers, fmt.Sprintf("phase2 выполнена %d раз, ожидалось %d", p2.Load(), workers))
	check(early.Load() == 0, fmt.Sprintf("%d горутин начали phase2 до окончания phase1 у всех", early.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
