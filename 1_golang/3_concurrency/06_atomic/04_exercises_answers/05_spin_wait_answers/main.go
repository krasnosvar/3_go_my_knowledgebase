// Упражнение 05: ожидание флага с уступкой процессора (Решение).
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func waitFlag(flag *atomic.Bool, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for !flag.Load() {
		if time.Now().After(deadline) {
			return false
		}
		runtime.Gosched()
	}
	return true
}

func main() {
	fmt.Println("--- Проверка 05_spin_wait (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var flag atomic.Bool
	go func() {
		time.Sleep(10 * time.Millisecond)
		flag.Store(true)
	}()
	check(waitFlag(&flag, 300*time.Millisecond), "флаг выставлен через 10ms, ожидалось true")

	var never atomic.Bool
	start := time.Now()
	ok := waitFlag(&never, 20*time.Millisecond)
	elapsed := time.Since(start)
	check(!ok, "флаг не выставлялся — ожидалось false")
	check(elapsed >= 15*time.Millisecond, fmt.Sprintf("вернулись через %v, не дождавшись таймаута", elapsed))
	check(elapsed < 300*time.Millisecond, "ожидание не прервалось по таймауту")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
