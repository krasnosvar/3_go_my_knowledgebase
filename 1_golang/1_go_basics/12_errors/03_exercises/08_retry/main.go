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

// doWork выполняет работу. Если success == true, функция должна вернуть nil без ошибок.
// Внимание: остерегайтесь ловушки типизированного nil (typed nil interface)!
func doWork(success bool) error {
	var err *CustomError = nil
	if !success {
		err = &CustomError{Message: "something went wrong"}
		return err
	}
	// TODO: Верните результат так, чтобы при success == true возвращался чистый nil интерфейс.
	return err
}

func main() {
	fmt.Println("--- Проверка 08_typed_nil ---")

	err := doWork(true)

	// Если допущена ошибка типизированного nil, проверка err != nil вернет true
	t1 := err == nil

	if t1 {
		fmt.Println("Тест типизированного nil: PASS")
	} else {
		fmt.Printf("Тест типизированного nil: FAIL (err != nil, хотя success == true! err=%v)\n", err)
	}
}
