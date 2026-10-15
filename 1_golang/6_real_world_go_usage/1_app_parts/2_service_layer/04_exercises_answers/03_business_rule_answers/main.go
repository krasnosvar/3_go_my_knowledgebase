// Упражнение 03: бизнес-правило перевода денег (Решение).
package main

import (
	"errors"
	"fmt"
)

var (
	errInvalidAmount = errors.New("invalid amount")
	errNoAccount     = errors.New("no such account")
	errInsufficient  = errors.New("insufficient funds")
)

type bank struct{ balances map[string]int }

func (b *bank) Transfer(from, to string, amount int) error {
	if amount <= 0 || from == to {
		return errInvalidAmount
	}
	fb, ok1 := b.balances[from]
	_, ok2 := b.balances[to]
	if !ok1 || !ok2 {
		return errNoAccount
	}
	if fb < amount {
		return errInsufficient
	}
	b.balances[from] -= amount
	b.balances[to] += amount
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_business_rule (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b := &bank{balances: map[string]int{"a": 100, "b": 0}}
	check(b.Transfer("a", "b", 30) == nil, "перевод 30 из a в b успешен")
	check(b.balances["a"] == 70 && b.balances["b"] == 30, "балансы a=70, b=30")
	check(errors.Is(b.Transfer("a", "b", 1000), errInsufficient), "перевод 1000 → errInsufficient")
	check(b.balances["a"] == 70 && b.balances["b"] == 30, "после ошибки балансы не изменились")
	check(errors.Is(b.Transfer("a", "zzz", 1), errNoAccount), "неизвестный счёт → errNoAccount")
	check(errors.Is(b.Transfer("a", "b", 0), errInvalidAmount), "amount=0 → errInvalidAmount")
	check(errors.Is(b.Transfer("a", "a", 5), errInvalidAmount), "перевод самому себе → errInvalidAmount")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
