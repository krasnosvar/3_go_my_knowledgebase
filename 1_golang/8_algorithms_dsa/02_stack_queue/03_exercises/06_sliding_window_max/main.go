// Упражнение 06: максимум в скользящем окне (монотонная очередь).
//
// Задача: реализуйте maxSlidingWindow(nums, k) — максимумы всех окон длины k.
// Пример: [1 3 -1 -3 5 3 6 7], k=3 -> [3 3 5 5 6 7].
// Наивно — O(n*k). Цель — O(n): храните в деке ИНДЕКСЫ по убыванию значений;
//   - выбрасывайте из головы индексы, вышедшие за окно;
//   - выбрасывайте из хвоста индексы с меньшими значениями, чем текущее;
//   - голова дека — индекс максимума текущего окна.
//
// Если k <= 0 или k > len(nums) — верните nil.
//
// Запуск: go run ./03_exercises/06_sliding_window_max/main.go
package main

import "fmt"

// maxSlidingWindow возвращает максимумы окон длины k.
func maxSlidingWindow(nums []int, k int) []int {
	// TODO: монотонная очередь индексов.
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_sliding_window_max ---")

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
