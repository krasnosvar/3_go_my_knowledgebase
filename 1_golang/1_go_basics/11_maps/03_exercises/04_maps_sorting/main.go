package main

import (
	"fmt"
)

// getSortedValues принимает мапу и возвращает слайс ее значений,
// отсортированных по алфавитному порядку ключей.
func getSortedValues(m map[string]int) []int {
	// TODO: Извлеките ключи, отсортируйте их и соберите соответствующие значения.
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_maps_sorting ---")

	m := map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
	}

	res := getSortedValues(m)

	t1 := len(res) == 3
	t2 := false
	if len(res) == 3 {
		t2 = res[0] == 1 && res[1] == 2 && res[2] == 3 // Соответствует порядку alice, bob, charlie
	}

	if t1 && t2 {
		fmt.Println("Тест сортировки мапы: PASS")
	} else {
		fmt.Printf("Тест сортировки мапы: FAIL (получено: %v)\n", res)
	}
}
