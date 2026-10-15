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

func doWork(success bool) error {
	if !success {
		return &CustomError{Message: "something went wrong"}
	}
	return nil // Возвращаем чистый untyped nil
}

func main() {
	fmt.Println("--- Проверка 08_typed_nil (Решение) ---")

	err := doWork(true)

	t1 := err == nil

	if t1 {
		fmt.Println("Тест типизированного nil: PASS")
	} else {
		fmt.Printf("Тест типизированного nil: FAIL (err != nil, хотя success == true! err=%v)\n", err)
	}
}
