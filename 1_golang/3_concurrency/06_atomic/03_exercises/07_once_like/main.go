// Упражнение 07: однократный запуск на CAS.
//
// Задача: OnceFlag.Do(f) вызывает f только при первом вызове Do, используя
// done.CompareAndSwap(0, 1); остальные вызовы сразу возвращаются.
//
// Отличие от sync.Once: проигравшие гонку НЕ ждут завершения f. Если
// вызывающему важно, что инициализация уже закончена, нужен sync.Once
// (внутри у него atomic + mutex именно ради ожидания).
//
// Запуск: go run ./03_exercises/07_once_like/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// OnceFlag выполняет функцию не более одного раза.
type OnceFlag struct {
	done atomic.Uint32
}

// Do вызывает f, только если это первый вызов Do.
func (o *OnceFlag) Do(f func()) {
	// TODO: if o.done.CompareAndSwap(0, 1) { f() }
}

func main() {
	fmt.Println("--- Проверка 07_once_like ---")

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
