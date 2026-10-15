// Упражнение 05: мутация состояния через метод.
//
// Задача: в Account есть баг — Deposit объявлен с value receiver, поэтому
// пополнение применяется к копии и теряется. Исправьте Deposit и реализуйте
// Withdraw:
//   - Deposit(n) увеличивает баланс;
//   - Withdraw(n) списывает n и возвращает true, если денег хватает;
//     иначе баланс не меняется и возвращается false.
//
// Запуск: go run ./03_exercises/05_mutation/main.go
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

// Deposit пополняет счёт.
// TODO: баг — value receiver меняет копию. Сделайте receiver указателем.
func (a Account) Deposit(n int) {
	a.balance += n
}

// Withdraw списывает n, если хватает средств.
func (a *Account) Withdraw(n int) bool {
	// TODO: если a.balance >= n — спишите и верните true, иначе false.
	return false
}

func main() {
	fmt.Println("--- Проверка 05_mutation ---")

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
