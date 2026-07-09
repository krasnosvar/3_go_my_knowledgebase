package main

import (
	"fmt"
	"slices"
)

func getSortedValues(m map[string]int) []int {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	vals := make([]int, 0, len(m))
	for _, k := range keys {
		vals = append(vals, m[k])
	}
	return vals
}

func main() {
	fmt.Println("--- Проверка 04_maps_sorting (Решение) ---")

	m := map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
	}

	res := getSortedValues(m)

	t1 := len(res) == 3
	t2 := false
	if len(res) == 3 {
		t2 = res[0] == 1 && res[1] == 2 && res[2] == 3
	}

	if t1 && t2 {
		fmt.Println("Тест сортировки мапы: PASS")
	} else {
		fmt.Printf("Тест сортировки мапы: FAIL (получено: %v)\n", res)
	}
}
