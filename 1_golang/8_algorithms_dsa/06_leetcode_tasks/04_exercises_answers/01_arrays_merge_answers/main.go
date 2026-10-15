// Упражнение 01: слияние отсортированных массивов in-place (LeetCode 88) (Решение).
package main

import "fmt"

// merge сливает nums2 в nums1 на месте.
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	fmt.Println("--- Проверка 01_arrays_merge (Решение) ---")

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
