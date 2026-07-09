package main

import (
	"fmt"
)

// UserID - именованный тип на базе int64.
// Помогает защитить API от перепутывания ID пользователей с другими ID.
type UserID int64

// Money - тип для денежных расчетов (в копейках/центах).
// ВАЖНО: Хранение денег в float64 на бэкенде — частая ошибка, которая
// приводит к потере копеек из-за погрешностей округления.
type Money int64

// Format превращает копейки в читаемый формат рублей (например, 10050 -> "100.50 руб.")
func (m Money) Format() string {
	return fmt.Sprintf("%.2f руб.", float64(m)/100)
}

func main() {
	var uid UserID = 100234

	// Стоимость товаров в копейках
	var item1 Money = 15000 // 150 руб. 00 коп.
	var item2 Money = 45050 // 450 руб. 50 коп.

	// Выполняем сложение. Поскольку оба операнда имеют тип Money, сложение корректно.
	total := item1 + item2

	fmt.Println("--- Бэкенд заказов ---")
	fmt.Printf("Пользователь ID: %d совершил покупку\n", uid)
	fmt.Printf("Товар 1: %s\n", item1.Format())
	fmt.Printf("Товар 2: %s\n", item2.Format())
	fmt.Printf("Итоговая сумма: %s\n", total.Format())
}
