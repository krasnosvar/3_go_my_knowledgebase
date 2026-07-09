package main

import (
	"fmt"
)

// sumEvenUntilLimit суммирует только ЧЕТНЫЕ числа из слайса nums.
// Если текущая накопленная сумма превышает limit, цикл должен прерваться.
// Используйте continue для пропуска нечетных чисел и break для выхода при превышении limit.
func sumEvenUntilLimit(nums []int, limit int) int {
	// TODO: Реализуйте логику с break и continue.
	return 0
}

func main() {
	fmt.Println("--- Проверка 05_break_continue ---")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// Только четные: 2, 4, 6, 8.
	// Лимит 10: 2 + 4 = 6 (ок), 6 + 6 = 12 (больше 10 -> break). Результат: 12.
	r1 := sumEvenUntilLimit(s, 10)
	t1 := r1 == 12

	// Лимит 100: 2+4+6+8 = 20 (не превышает 100). Результат: 20.
	r2 := sumEvenUntilLimit(s, 100)
	t2 := r2 == 20

	if t1 && t2 {
		fmt.Println("Тест break/continue: PASS")
	} else {
		fmt.Printf("Тест break/continue: FAIL (r1=%d, r2=%d)\n", r1, r2)
	}
}
