package main

import (
	"fmt"
)

// modifySliceElements изменяет первый элемент слайса на 100.
// Поскольку слайс под капотом хранит указатель на массив, это изменение отразится в вызывающей функции.
func modifySliceElements(s []int) {
	// TODO: Присвойте первому элементу слайса s (индекс 0) значение 100.
}

// appendByValue пытается добавить число 42 в слайс, переданный по значению.
// Поскольку заголовок слайса (len, cap) копируется, это изменение НЕ отразится снаружи.
func appendByValue(s []int) {
	s = append(s, 42)
}

// appendByPointer добавляет число 42 в слайс через указатель на слайс.
// Это изменение отразится снаружи.
func appendByPointer(s *[]int) {
	// TODO: Разыменуйте указатель и добавьте число 42 в слайс с помощью append.
}

func main() {
	fmt.Println("--- Проверка 07_slice_pointer ---")

	// 1. Тест изменения элементов
	nums1 := []int{1, 2, 3}
	modifySliceElements(nums1)
	t1 := nums1[0] == 100

	// 2. Тест добавления по значению (не должно сработать)
	nums2 := []int{1, 2, 3}
	appendByValue(nums2)
	t2 := len(nums2) == 3 // длина не должна измениться снаружи!

	// 3. Тест добавления по указателю
	nums3 := []int{1, 2, 3}
	appendByPointer(&nums3)
	t3 := len(nums3) == 4 && nums3[3] == 42

	if t1 && t2 && t3 {
		fmt.Println("Тест работы со слайсами: PASS")
	} else {
		fmt.Printf("Тест работы со слайсами: FAIL (t1=%t, t2=%t, t3=%t)\n", t1, t2, t3)
	}
}
