// Упражнение 08: барьер между фазами (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func runPhases(workers int, phase1, phase2 func(id int)) {
	var phase1Done, all sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < workers; i++ {
		phase1Done.Add(1)
		all.Add(1)
		go func(id int) {
			defer all.Done()
			phase1(id)
			phase1Done.Done()
			<-start
			phase2(id)
		}(i)
	}
	phase1Done.Wait()
	close(start)
	all.Wait()
}

func main() {
	fmt.Println("--- Проверка 08_barrier (Решение) ---")

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
