package main

import (
	"fmt"
)

// MyError — кастомная структура ошибки, содержащая Code (int) и Message (string).
// TODO: Объявите структуру и реализуйте метод Error() string.

func main() {
	fmt.Println("--- Проверка 05_custom_type ---")

	// После того, как вы реализуете MyError, раскомментируйте код проверки ниже:
	/*
		var err error = &MyError{Code: 404, Message: "Not Found"}

		t1 := err.Error() == "Error 404: Not Found"

		if t1 {
			fmt.Println("Тест MyError: PASS")
		} else {
			fmt.Printf("Тест MyError: FAIL (получено %q)\n", err.Error())
		}
	*/

	fmt.Println("Реализуйте MyError с методом Error() и раскомментируйте тест в main!")
}
