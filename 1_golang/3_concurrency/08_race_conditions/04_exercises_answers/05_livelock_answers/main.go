// Упражнение 05: livelock и случайная задержка (Решение).
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func acquireBoth(first, second *sync.Mutex) bool {
	for {
		first.Lock()
		if second.TryLock() {
			return true
		}
		first.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	}
}

func main() {
	fmt.Println("--- Проверка 05_livelock (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var a, b sync.Mutex
	var success atomic.Int64
	done := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		worker := func(x, y *sync.Mutex) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				if acquireBoth(x, y) {
					success.Add(1)
					y.Unlock()
					x.Unlock()
				}
			}
		}
		wg.Add(2)
		go worker(&a, &b)
		go worker(&b, &a)
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		check(false, "горутины не завершились — livelock или deadlock")
	}
	check(success.Load() == 400, fmt.Sprintf("успешных захватов %d, ожидалось 400", success.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
