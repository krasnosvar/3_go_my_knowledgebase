// Упражнение 02: инвариант баланса под мьютексом (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	fmt.Println("--- Проверка 02_fix_with_mutex (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var acc Account
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(10)
		}()
	}
	wg.Wait()
	check(acc.Balance() == 1000, fmt.Sprintf("после пополнений баланс = %d, ожидалось 1000", acc.Balance()))

	var ok atomic.Int64
	for i := 0; i < 150; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if acc.Withdraw(10) {
				ok.Add(1)
			}
		}()
	}
	wg.Wait()
	check(ok.Load() == 100, fmt.Sprintf("успешных списаний %d, ожидалось ровно 100", ok.Load()))
	check(acc.Balance() == 0, fmt.Sprintf("итоговый баланс = %d, ожидалось 0 (не меньше!)", acc.Balance()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
