package main

import (
	"fmt"
	"math"
)

// avg возвращает среднее арифметическое двух целых чисел.
// Например, avg(3, 4) должно вернуть 3.5, а не 3.0.
func avg(a int, b int) float64 {
	// Приводим хотя бы один (или оба) операнда к float64 перед делением.
	// Если сделать float64((a + b) / 2), то деление произойдет как целочисленное
	// (дробная часть отсечется), и только потом результат приведется к float64.
	return float64(a+b) / 2.0
}

func main() {
	fmt.Println("--- Проверка 05_type_conversion (Решение) ---")

	tests := []struct {
		a    int
		b    int
		want float64
	}{
		{3, 4, 3.5},
		{10, 20, 15.0},
		{0, 0, 0.0},
	}

	allPassed := true
	for i, tt := range tests {
		got := avg(tt.a, tt.b)
		// Для сравнения float используем epsilon-разницу
		if math.Abs(got-tt.want) < 1e-9 {
			fmt.Printf("Тест %d (avg от %d и %d): PASS\n", i+1, tt.a, tt.b)
		} else {
			fmt.Printf("Тест %d (avg от %d и %d): FAIL (ожидалось %.2f, получено %.2f)\n",
				i+1, tt.a, tt.b, tt.want, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
