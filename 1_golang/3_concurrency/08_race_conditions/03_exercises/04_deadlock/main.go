// Упражнение 04: deadlock при переводе между счетами.
//
// Задача: transfer(from, to, amount) блокирует ОБА счёта и переводит
// деньги. Наивная версия захватывает from.mu, потом to.mu — и если
// параллельно идут переводы a→b и b→a, горутины ждут друг друга вечно
// (deadlock: каждая держит один мьютекс и ждёт второй).
//
// Исправление — единый порядок захвата: всегда сначала мьютекс счёта
// с меньшим id, потом с большим. Тест гоняет встречные переводы и
// падает с FAIL, если всё зависло.
//
// Запуск: go run ./03_exercises/04_deadlock/main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Acc — счёт с уникальным id.
type Acc struct {
	id      int
	mu      sync.Mutex
	balance int
}

// transfer переводит amount со счёта from на счёт to.
func transfer(from, to *Acc, amount int) {
	// TODO: определить first/second по id, Lock в этом порядке, перевести.
}

func main() {
	fmt.Println("--- Проверка 04_deadlock ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := &Acc{id: 1, balance: 1000}
	b := &Acc{id: 2, balance: 1000}

	done := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				transfer(a, b, 1)
			}()
			go func() {
				defer wg.Done()
				transfer(b, a, 2)
			}()
		}
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		a.mu.Lock()
		b.mu.Lock()
		check(a.balance == 1100, fmt.Sprintf("баланс a = %d, ожидалось 1100", a.balance))
		check(b.balance == 900, fmt.Sprintf("баланс b = %d, ожидалось 900", b.balance))
		b.mu.Unlock()
		a.mu.Unlock()
	case <-time.After(time.Second):
		check(false, "переводы зависли — deadlock")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
