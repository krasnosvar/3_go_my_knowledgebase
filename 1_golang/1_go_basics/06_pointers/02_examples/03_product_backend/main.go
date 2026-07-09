package main

import (
	"fmt"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

// Потенциально опасная функция. Она возвращает типизированный указатель *CustomError.
// Если ошибки нет, она возвращает nil, приведенный к типу *CustomError.
func processOrder(amount int) *CustomError {
	if amount < 0 {
		return &CustomError{Message: "отрицательная сумма заказа"}
	}
	// Ошибки нет, возвращаем nil типа *CustomError
	var err *CustomError = nil
	return err
}

func main() {
	fmt.Println("--- Бэкенд заказов: ловушка типизированного nil в интерфейсе ---")

	// Мы присваиваем возвращаемое значение переменной типа error (интерфейс!)
	var err error = processOrder(100)

	// ВНИМАНИЕ: Ожидается, что err равен nil, так как сумма заказа корректна (100).
	// Однако интерфейс error хранит структуру: (type: *CustomError, value: nil).
	// Поскольку поле type заполнено, сам интерфейс не равен nil!
	if err != nil {
		fmt.Printf("ВНИМАНИЕ (Баг!): Обнаружена ошибка, хотя заказ корректен: %v\n", err)
		fmt.Printf("Тип ошибки в интерфейсе: %T, Значение: %v\n", err, err)
	} else {
		fmt.Println("Все отлично! Ошибок нет.")
	}

	fmt.Println("\n--- Как исправить этот баг ---")
	// Функция должна явно возвращать nil как нетипизированный nil-интерфейс error.
	var fixedErr error = processOrderFixed(100)
	if fixedErr != nil {
		fmt.Println("Ошибка:", fixedErr)
	} else {
		fmt.Println("Все отлично! Ошибок нет (баг исправлен).")
	}
}

// Исправленная версия функции: возвращает сразу интерфейсный тип error
func processOrderFixed(amount int) error {
	if amount < 0 {
		return &CustomError{Message: "отрицательная сумма заказа"}
	}
	// Возвращаем чистый nil
	return nil
}
