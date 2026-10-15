// Упражнение 07: однократный запуск на CAS (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type OnceFlag struct {
	done atomic.Uint32
}

func (o *OnceFlag) Do(f func()) {
	if o.done.CompareAndSwap(0, 1) {
		f()
	}
}

func main() {
	fmt.Println("--- Проверка 07_once_like (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var o OnceFlag
	var calls atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.Do(func() { calls.Add(1) })
		}()
	}
	wg.Wait()
	check(calls.Load() == 1, fmt.Sprintf("f вызвана %d раз, ожидалось 1", calls.Load()))

	o.Do(func() { calls.Add(1) })
	check(calls.Load() == 1, "повторный Do не должен вызывать f")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
