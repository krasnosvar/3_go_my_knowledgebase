// Упражнение 06: максимум в скользящем окне (монотонная очередь) (Решение).
package main

import "fmt"

// maxSlidingWindow возвращает максимумы окон длины k.
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return nil
	}
	var deque []int // индексы, значения nums по ним убывают
	res := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= v {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

func main() {
	fmt.Println("--- Проверка 06_sliding_window_max (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	check(fmt.Sprint(got) == "[3 3 5 5 6 7]", fmt.Sprintf("ожидалось [3 3 5 5 6 7], получено %v", got))

	got = maxSlidingWindow([]int{9, 8, 7}, 1)
	check(fmt.Sprint(got) == "[9 8 7]", fmt.Sprintf("k=1: ожидалось [9 8 7], получено %v", got))

	check(maxSlidingWindow([]int{1, 2}, 5) == nil, "k > len(nums): ожидался nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
