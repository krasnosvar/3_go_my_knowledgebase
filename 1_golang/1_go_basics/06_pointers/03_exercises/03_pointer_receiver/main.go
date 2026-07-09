package main

import (
	"fmt"
)

type Account struct {
	Balance int
}

// TODO: Определите метод Deposit(amount int) с получателем-указателем (*Account),
// который увеличивает Balance на переданную сумму.

func main() {
	fmt.Println("--- Проверка 03_pointer_receiver ---")

	// ВНИМАНИЕ: Для прохождения теста объявите метод Deposit выше main
	// и раскомментируйте код ниже для проверки

	/*
		acc := Account{Balance: 100}
		acc.Deposit(50)

		if acc.Balance == 150 {
			fmt.Println("Тест Deposit: PASS")
		} else {
			fmt.Printf("Тест Deposit: FAIL (ожидалось 150, получено %d - возможно, метод объявлен без указателя)\n", acc.Balance)
		}
	*/

	fmt.Println("Раскомментируйте блок проверки в main")
}
