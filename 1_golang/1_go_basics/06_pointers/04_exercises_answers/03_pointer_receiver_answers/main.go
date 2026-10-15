package main

import (
	"fmt"
)

type Account struct {
	Balance int
}

// Deposit увеличивает Balance на переданную сумму
func (acc *Account) Deposit(amount int) {
	if acc != nil {
		acc.Balance += amount
	}
}

func main() {
	fmt.Println("--- Проверка 03_pointer_receiver (Решение) ---")

	acc := Account{Balance: 100}
	acc.Deposit(50)

	if acc.Balance == 150 {
		fmt.Println("Тест Deposit: PASS")
	} else {
		fmt.Printf("Тест Deposit: FAIL (ожидалось 150, получено %d)\n", acc.Balance)
	}
}
