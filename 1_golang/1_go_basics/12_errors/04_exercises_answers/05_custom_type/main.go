package main

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {
	fmt.Println("--- Проверка 05_custom_type (Решение) ---")

	var err error = &MyError{Code: 404, Message: "Not Found"}

	t1 := err.Error() == "Error 404: Not Found"

	if t1 {
		fmt.Println("Тест MyError: PASS")
	} else {
		fmt.Printf("Тест MyError: FAIL (получено %q)\n", err.Error())
	}
}
