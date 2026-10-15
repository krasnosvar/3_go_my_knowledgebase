// Упражнение 04: ленивая инициализация через sync.Once.
//
// Задача: Lazy.Get() при первом вызове вычисляет значение через l.init()
// и запоминает его; все последующие (в том числе конкурентные) вызовы
// возвращают сохранённое значение, а init больше не вызывается.
// Используйте l.once.Do(...).
//
// Запуск: go run ./03_exercises/04_once/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Lazy вычисляет значение один раз при первом обращении.
type Lazy struct {
	once sync.Once
	val  int
	init func() int
}

// Get возвращает значение, инициализируя его ровно один раз.
func (l *Lazy) Get() int {
	// TODO: l.once.Do(func() { l.val = l.init() }); return l.val
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_once ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var calls atomic.Int64
	l := &Lazy{init: func() int {
		calls.Add(1)
		time.Sleep(5 * time.Millisecond)
		return 42
	}}

	var wg sync.WaitGroup
	var wrong atomic.Int64
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if l.Get() != 42 {
				wrong.Add(1)
			}
		}()
	}
	wg.Wait()
	check(calls.Load() == 1, fmt.Sprintf("init вызвана %d раз, ожидалось 1", calls.Load()))
	check(wrong.Load() == 0, fmt.Sprintf("%d вызовов Get вернули не 42", wrong.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
