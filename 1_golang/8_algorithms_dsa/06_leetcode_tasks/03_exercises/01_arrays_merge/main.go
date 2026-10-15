// Упражнение 01: слияние отсортированных массивов in-place (LeetCode 88).
//
// Задача: реализуйте merge(nums1, m, nums2, n). В nums1 первые m элементов
// отсортированы, в хвосте n «пустых» нулей; nums2 содержит n отсортированных
// элементов. Слейте nums2 в nums1 НА МЕСТЕ так, чтобы nums1 был отсортирован.
// Трюк: заполняйте nums1 с КОНЦА тремя указателями (i=m-1, j=n-1, k=m+n-1) —
// так ничего не затирается и не нужен буфер. O(m+n) времени, O(1) памяти.
//
// Запуск: go run ./03_exercises/01_arrays_merge/main.go
package main

import "fmt"

// merge сливает nums2 в nums1 на месте.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// TODO: три указателя, заполнение с конца.
}

func main() {
	fmt.Println("--- Проверка 01_arrays_merge ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	check(fmt.Sprint(a) == "[1 2 2 3 5 6]", fmt.Sprintf("ожидалось [1 2 2 3 5 6], получено %v", a))

	b := []int{0}
	merge(b, 0, []int{1}, 1)
	check(fmt.Sprint(b) == "[1]", fmt.Sprintf("m=0: ожидалось [1], получено %v", b))

	c := []int{4, 5, 6, 0, 0}
	merge(c, 3, []int{1, 2}, 2)
	check(fmt.Sprint(c) == "[1 2 4 5 6]", fmt.Sprintf("ожидалось [1 2 4 5 6], получено %v", c))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
