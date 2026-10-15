// Упражнение 05: ожидание флага с уступкой процессора.
//
// Задача: waitFlag(flag, timeout) крутится в цикле, пока flag не станет
// true, и возвращает true; если за timeout флаг не выставлен — false.
// В теле цикла вызывайте runtime.Gosched(), чтобы уступать процессор
// другим горутинам, а не «жечь» его впустую.
//
// Spin-wait уместен только для очень коротких ожиданий; для долгих —
// каналы, sync.Cond или context.
//
// Запуск: go run ./03_exercises/05_spin_wait/main.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// waitFlag ждёт flag == true не дольше timeout.
func waitFlag(flag *atomic.Bool, timeout time.Duration) bool {
	// TODO: deadline := time.Now().Add(timeout); цикл с flag.Load() и runtime.Gosched().
	return false
}

func main() {
	fmt.Println("--- Проверка 05_spin_wait ---")

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
