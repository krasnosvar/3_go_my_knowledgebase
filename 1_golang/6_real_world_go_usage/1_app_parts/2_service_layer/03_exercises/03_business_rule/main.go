// Упражнение 03: бизнес-правило перевода денег.
//
// Задача: реализуйте (b *bank) Transfer(from, to string, amount int) error:
//   - amount <= 0 → errInvalidAmount;
//   - неизвестный счёт (from или to) → errNoAccount;
//   - from == to → errInvalidAmount;
//   - недостаточно средств → errInsufficient, балансы НЕ меняются;
//   - иначе списать с from и зачислить на to.
//
// Все проверки — до любых изменений (никаких частичных обновлений).
//
// Запуск: go run ./03_exercises/03_business_rule
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
	// TODO: сначала все проверки, потом изменение балансов.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_business_rule ---")

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
