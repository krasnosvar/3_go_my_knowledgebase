package main

import (
	"fmt"
)

// getFirstTwoElements возвращает первые 2 элемента из слайса hugeSlice.
// Чтобы избежать удержания огромного массива в памяти, функция должна скопировать
// эти 2 элемента в новый слайс с помощью copy() и вернуть его.
func getFirstTwoElements(hugeSlice []int) []int {
	// TODO: Реализуйте копирование во избежание утечки памяти.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_memory_retention ---")

	huge := []int{10, 20, 30, 40, 50, 60}
	clip := getFirstTwoElements(huge)

	t1 := len(clip) == 2 && clip[0] == 10 && clip[1] == 20

	// Проверяем, что clip указывает на другой backing array.
	// Если мы поменяем clip[0], huge[0] НЕ должен измениться.
	clip[0] = 99
	t2 := huge[0] == 10

	if t1 && t2 {
		fmt.Println("Тест утечки памяти: PASS")
	} else {
		fmt.Printf("Тест утечки памяти: FAIL (huge[0]=%d, clip=%v)\n", huge[0], clip)
	}
}
