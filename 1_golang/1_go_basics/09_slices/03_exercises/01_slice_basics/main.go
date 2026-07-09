package main

import (
	"fmt"
)

// createIntSlice создает слайс целых чисел длины length и емкости capacity с помощью make,
// и заполняет его числами от 0 до length-1.
func createIntSlice(length, capacity int) []int {
	// TODO: Реализуйте создание и заполнение слайса.
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_slice_basics ---")

	s := createIntSlice(3, 5)

	t1 := s != nil
	t2 := len(s) == 3
	t3 := cap(s) == 5
	t4 := false
	if len(s) == 3 {
		t4 = s[0] == 0 && s[1] == 1 && s[2] == 2
	}

	if t1 && t2 && t3 && t4 {
		fmt.Println("Тест createIntSlice: PASS")
	} else {
		fmt.Printf("Тест createIntSlice: FAIL (len=%d, cap=%d, val=%v)\n", len(s), cap(s), s)
	}
}
