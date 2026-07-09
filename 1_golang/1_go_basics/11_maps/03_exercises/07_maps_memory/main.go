package main

import (
	"fmt"
)

// shrinkMap проверяет, уменьшилось ли число элементов мапы ниже порога threshold.
// Если да, функция создает новую мапу, копирует туда элементы и возвращает ее,
// позволяя сборщику мусора освободить старые неиспользуемые бакеты.
// В противном случае возвращает исходную мапу.
func shrinkMap(m map[int]string, threshold int) map[int]string {
	// TODO: Реализуйте усадку мапы при падении числа элементов ниже порога.
	return m
}

func main() {
	fmt.Println("--- Проверка 07_maps_memory ---")

	m := make(map[int]string)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	// Удаляем 950 элементов
	for i := 0; i < 950; i++ {
		delete(m, i)
	}

	// Порог усадки — 100 элементов
	shrunk := shrinkMap(m, 100)

	t1 := len(shrunk) == 50

	// Проверяем, что это физически разные мапы (адресация/изменения)
	shrunk[9999] = "new"
	_, foundOld := m[9999]

	if t1 && !foundOld {
		fmt.Println("Тест shrinkMap: PASS")
	} else {
		fmt.Printf("Тест shrinkMap: FAIL (len=%d, foundOld=%t)\n", len(shrunk), foundOld)
	}
}
