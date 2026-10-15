// Упражнение 08: корректная остановка нескольких горутин.
//
// Задача: startWorkers(n, stop, wg, ticks) запускает n горутин. Каждая в
// цикле увеличивает ticks на 1 и спит 1ms, пока не закрыт stop. Каждая
// горутина должна быть учтена в wg (wg.Add до запуска, defer wg.Done()).
//
// После close(stop) вызывающий код делает wg.Wait() и уверен, что все
// горутины завершились. Проверку stop делайте через select с default
// или через select с time.After(time.Millisecond).
//
// Запуск: go run ./03_exercises/08_graceful_stop/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// startWorkers запускает n горутин, работающих до закрытия stop.
func startWorkers(n int, stop <-chan struct{}, wg *sync.WaitGroup, ticks *atomic.Int64) {
	// TODO: запустите n горутин, учтите их в wg, в цикле увеличивайте ticks
	// и выходите при закрытии stop.
}

func main() {
	fmt.Println("--- Проверка 08_graceful_stop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var (
		wg    sync.WaitGroup
		ticks atomic.Int64
	)
	stop := make(chan struct{})
	startWorkers(3, stop, &wg, &ticks)

	time.Sleep(30 * time.Millisecond)
	close(stop)

	waited := make(chan struct{})
	go func() {
		wg.Wait()
		close(waited)
	}()
	select {
	case <-waited:
	case <-time.After(300 * time.Millisecond):
		check(false, "горутины не завершились после close(stop)")
	}

	check(ticks.Load() >= 3, fmt.Sprintf("ticks = %d: горутины не работали", ticks.Load()))
	after := ticks.Load()
	time.Sleep(10 * time.Millisecond)
	check(ticks.Load() == after, "после остановки ticks продолжает расти")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
