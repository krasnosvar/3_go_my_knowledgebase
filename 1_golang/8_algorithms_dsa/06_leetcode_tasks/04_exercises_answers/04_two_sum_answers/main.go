// Упражнение 04: Two Sum через хеш-таблицу (LeetCode 1) (Решение).
package main

import "fmt"

// twoSum возвращает индексы двух элементов с суммой target.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for j, x := range nums {
		if i, ok := seen[target-x]; ok {
			return []int{i, j}
		}
		seen[x] = j
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_two_sum (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := twoSum([]int{2, 7, 11, 15}, 9)
	check(fmt.Sprint(got) == "[0 1]", fmt.Sprintf("ожидалось [0 1], получено %v", got))

	got = twoSum([]int{3, 2, 4}, 6)
	check(fmt.Sprint(got) == "[1 2]", fmt.Sprintf("ожидалось [1 2] (нельзя брать один элемент дважды), получено %v", got))

	got = twoSum([]int{3, 3}, 6)
	check(fmt.Sprint(got) == "[0 1]", fmt.Sprintf("дубликаты: ожидалось [0 1], получено %v", got))

	check(twoSum([]int{1, 2}, 10) == nil, "нет ответа: ожидался nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
