package main

import (
	"fmt"
)

type Account struct {
	Balance int
}

// ⏩ Забегание вперёд: структуры подробно разбираются в теме 08_structs,
// а методы — в 2_oop/01_methods. Здесь важно только одно: получатель метода
// ведёт себя как обычный параметр — значение копируется, указатель — нет.

// Deposit увеличивает Balance на переданную сумму.
// TODO: Сейчас метод объявлен с получателем-ЗНАЧЕНИЕМ (Account): он меняет копию,
// и баланс исходного счёта не растёт. Исправьте получатель на указатель (*Account).
func (a Account) Deposit(amount int) {
	a.Balance += amount
}

func main() {
	fmt.Println("--- Проверка 03_pointer_receiver ---")

	acc := Account{Balance: 100}
	acc.Deposit(50)

	if acc.Balance == 150 {
		fmt.Println("Тест Deposit: PASS")
	} else {
		fmt.Printf("Тест Deposit: FAIL (ожидалось 150, получено %d - возможно, метод объявлен без указателя)\n", acc.Balance)
	}
}
