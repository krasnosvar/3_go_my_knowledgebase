package main

import (
	"fmt"
)

func shrinkMap(m map[int]string, threshold int) map[int]string {
	if len(m) < threshold {
		newMap := make(map[int]string, len(m))
		for k, v := range m {
			newMap[k] = v
		}
		return newMap
	}
	return m
}

func main() {
	fmt.Println("--- Проверка 07_maps_memory (Решение) ---")

	m := make(map[int]string)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	for i := 0; i < 950; i++ {
		delete(m, i)
	}

	shrunk := shrinkMap(m, 100)

	t1 := len(shrunk) == 50

	shrunk[9999] = "new"
	_, foundOld := m[9999]

	if t1 && !foundOld {
		fmt.Println("Тест shrinkMap: PASS")
	} else {
		fmt.Printf("Тест shrinkMap: FAIL (len=%d, foundOld=%t)\n", len(shrunk), foundOld)
	}
}
