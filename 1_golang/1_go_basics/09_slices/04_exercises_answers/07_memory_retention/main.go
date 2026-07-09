package main

import (
	"fmt"
)

// getFirstTwoElements возвращает первые 2 элемента без удержания исходного массива в памяти
func getFirstTwoElements(hugeSlice []int) []int {
	if len(hugeSlice) < 2 {
		return nil
	}
	clip := make([]int, 2)
	copy(clip, hugeSlice[:2])
	return clip
}

func main() {
	fmt.Println("--- Проверка 07_memory_retention (Решение) ---")

	huge := []int{10, 20, 30, 40, 50, 60}
	clip := getFirstTwoElements(huge)

	t1 := len(clip) == 2 && clip[0] == 10 && clip[1] == 20

	clip[0] = 99
	t2 := huge[0] == 10

	if t1 && t2 {
		fmt.Println("Тест утечки памяти: PASS")
	} else {
		fmt.Printf("Тест утечки памяти: FAIL (huge[0]=%d, clip=%v)\n", huge[0], clip)
	}
}
