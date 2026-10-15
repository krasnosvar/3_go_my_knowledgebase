package main

import (
	"errors"
	"fmt"
)

// divide возвращает частное и остаток
func divide(a, b int) (quotient int, remainder int, err error) {
	if b == 0 {
		return 0, 0, errors.New("деление на ноль")
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, nil
}

func main() {
	fmt.Println("--- Проверка 02_multiple_return (Решение) ---")

	q1, r1, err1 := divide(10, 3)
	_, _, err2 := divide(10, 0)

	t1 := q1 == 3 && r1 == 1 && err1 == nil
	t2 := err2 != nil

	if t1 && t2 {
		fmt.Println("Тест divide: PASS")
	} else {
		fmt.Printf("Тест divide: FAIL (q1=%d, r1=%d, err1=%v, err2=%v)\n", q1, r1, err1, err2)
	}
}
