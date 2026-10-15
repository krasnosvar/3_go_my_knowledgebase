// Упражнение 05: мутация состояния через метод (Решение).
package main

import (
	"fmt"
)

// Account — банковский счёт.
type Account struct {
	balance int
}

// Balance возвращает текущий баланс.
func (a Account) Balance() int { return a.balance }

func (a *Account) Deposit(n int) {
	a.balance += n
}

func (a *Account) Withdraw(n int) bool {
	if a.balance < n {
		return false
	}
	a.balance -= n
	return true
}

func main() {
	fmt.Println("--- Проверка 05_mutation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	acc := &Account{}
	acc.Deposit(100)
	check(acc.Balance() == 100, fmt.Sprintf("после Deposit(100): ожидалось 100, получено %d", acc.Balance()))

	ok := acc.Withdraw(30)
	check(ok && acc.Balance() == 70, fmt.Sprintf("Withdraw(30): ожидалось ok=true и 70, получено ok=%t и %d", ok, acc.Balance()))

	ok = acc.Withdraw(500)
	check(!ok && acc.Balance() == 70, fmt.Sprintf("Withdraw(500) при 70: ожидалось ok=false и 70, получено ok=%t и %d", ok, acc.Balance()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
