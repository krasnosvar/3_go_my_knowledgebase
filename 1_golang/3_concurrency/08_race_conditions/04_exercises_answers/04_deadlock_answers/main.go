// Упражнение 04: deadlock при переводе между счетами (Решение).
package main

import (
	"fmt"
	"sync"
	"time"
)

type Acc struct {
	id      int
	mu      sync.Mutex
	balance int
}

func transfer(from, to *Acc, amount int) {
	first, second := from, to
	if second.id < first.id {
		first, second = second, first
	}
	first.mu.Lock()
	defer first.mu.Unlock()
	second.mu.Lock()
	defer second.mu.Unlock()
	from.balance -= amount
	to.balance += amount
}

func main() {
	fmt.Println("--- Проверка 04_deadlock (Решение) ---")

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
