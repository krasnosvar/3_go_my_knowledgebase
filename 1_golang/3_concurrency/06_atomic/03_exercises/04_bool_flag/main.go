// Упражнение 04: флаг закрытия на atomic.Bool.
//
// Задача: Shutdown.Close() должна вернуть true ровно у ОДНОГО вызывающего
// (того, кто первым перевёл флаг false → true), у остальных — false.
// IsClosed() сообщает текущее состояние.
//
// Используйте s.closed.CompareAndSwap(false, true): «проверить и
// установить» в одной атомарной операции. Связка Load()+Store() здесь
// некорректна — между ними может вклиниться другая горутина.
//
// Запуск: go run ./03_exercises/04_bool_flag/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Shutdown — одноразовый флаг закрытия.
type Shutdown struct {
	closed atomic.Bool
}

// Close возвращает true только для первого вызова.
func (s *Shutdown) Close() bool {
	// TODO: return s.closed.CompareAndSwap(false, true)
	return false
}

// IsClosed сообщает, закрыт ли флаг.
func (s *Shutdown) IsClosed() bool {
	// TODO
	return false
}

func main() {
	fmt.Println("--- Проверка 04_bool_flag ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Shutdown
	check(!s.IsClosed(), "новый флаг не должен быть закрыт")

	var winners atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if s.Close() {
				winners.Add(1)
			}
		}()
	}
	wg.Wait()
	check(winners.Load() == 1, fmt.Sprintf("Close вернула true %d раз, ожидалось ровно 1", winners.Load()))
	check(s.IsClosed(), "после Close флаг должен быть закрыт")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
