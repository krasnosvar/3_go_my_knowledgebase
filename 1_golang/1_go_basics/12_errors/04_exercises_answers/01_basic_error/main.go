package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("--- Проверка 01_basic_error (Решение) ---")

	res1, err1 := divide(10, 2)
	res2, err2 := divide(10, 0)

	t1 := err1 == nil && res1 == 5
	t2 := err2 != nil && res2 == 0 && err2.Error() == "division by zero"

	if t1 && t2 {
		fmt.Println("Тест divide: PASS")
	} else {
		fmt.Printf("Тест divide: FAIL (res1=%d, err1=%v, res2=%d, err2=%v)\n", res1, err1, res2, err2)
	}
}
