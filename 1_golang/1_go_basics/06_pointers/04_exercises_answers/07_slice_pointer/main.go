package main

import (
	"fmt"
)

// modifySliceElements изменяет первый элемент слайса на 100
func modifySliceElements(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

// appendByValue пытается добавить число 42 в слайс, переданный по значению
func appendByValue(s []int) {
	s = append(s, 42)
}

// appendByPointer добавляет число 42 в слайс через указатель
func appendByPointer(s *[]int) {
	if s != nil {
		*s = append(*s, 42)
	}
}

func main() {
	fmt.Println("--- Проверка 07_slice_pointer (Решение) ---")

	// 1. Тест изменения элементов
	nums1 := []int{1, 2, 3}
	modifySliceElements(nums1)
	t1 := nums1[0] == 100

	// 2. Тест добавления по значению
	nums2 := []int{1, 2, 3}
	appendByValue(nums2)
	t2 := len(nums2) == 3

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
