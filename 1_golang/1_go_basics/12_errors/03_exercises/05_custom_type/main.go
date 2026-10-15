package main

import (
	"fmt"
)

// MyError — кастомная структура ошибки, содержащая Code (int) и Message (string).
// Метод Error() должен возвращать строку формата "Error <Code>: <Message>".
type MyError struct {
	Code    int
	Message string
}

// Error делает *MyError реализацией встроенного интерфейса error.
// TODO: верните строку формата "Error 404: Not Found" (подсказка: fmt.Sprintf).
func (e *MyError) Error() string {
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_custom_type ---")

	var err error = &MyError{Code: 404, Message: "Not Found"}

	t1 := err.Error() == "Error 404: Not Found"

	if t1 {
		fmt.Println("Тест MyError: PASS")
	} else {
		fmt.Printf("Тест MyError: FAIL (получено %q)\n", err.Error())
	}
}
