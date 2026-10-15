// Упражнение 02: инвариант баланса под мьютексом.
//
// Задача: Account с конкурентными Deposit/Withdraw. Без синхронизации
// возможна гонка «проверить-потом-списать»: две горутины одновременно
// видят достаточный баланс и обе списывают — баланс уходит в минус.
//
// Реализуйте под одним мьютексом (проверка и списание — в одной
// критической секции):
//   - Deposit(amount) — пополнение;
//   - Withdraw(amount) bool — списание, только если хватает средств;
//   - Balance() int — текущий баланс.
//
// Запуск: go run ./03_exercises/02_fix_with_mutex/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Account — банковский счёт с конкурентным доступом.
type Account struct {
	mu      sync.Mutex
	balance int
}

// Deposit пополняет счёт.
func (a *Account) Deposit(amount int) {
	// TODO
}

// Withdraw списывает amount, если хватает средств.
func (a *Account) Withdraw(amount int) bool {
	// TODO: проверка и списание под одним Lock.
	return false
}

// Balance возвращает текущий баланс.
func (a *Account) Balance() int {
	// TODO
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_fix_with_mutex ---")

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
