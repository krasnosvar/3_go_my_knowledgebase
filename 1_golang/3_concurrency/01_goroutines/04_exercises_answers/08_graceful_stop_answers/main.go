// Упражнение 08: корректная остановка нескольких горутин (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func startWorkers(n int, stop <-chan struct{}, wg *sync.WaitGroup, ticks *atomic.Int64) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				case <-time.After(time.Millisecond):
					ticks.Add(1)
				}
			}
		}()
	}
}

func main() {
	fmt.Println("--- Проверка 08_graceful_stop (Решение) ---")

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
