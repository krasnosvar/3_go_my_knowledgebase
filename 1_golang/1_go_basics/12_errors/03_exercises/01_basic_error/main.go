package main

import (
	"fmt"
)

// divide возвращает результат деления a на b.
// Если b == 0, возвращает 0 и ошибку (созданную через errors.New).
func divide(a, b int) (int, error) {
	// TODO: Реализуйте деление с возвратом ошибки при делении на 0.
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 01_basic_error ---")

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
