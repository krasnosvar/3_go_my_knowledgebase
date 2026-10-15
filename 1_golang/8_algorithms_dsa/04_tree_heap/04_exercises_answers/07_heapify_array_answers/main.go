// Упражнение 07: построение кучи из массива за O(n) (Решение).
package main

import (
	"fmt"
	"sort"
)

// isMinHeap проверяет свойство min-кучи: родитель <= детей.
func isMinHeap(a []int) bool {
	for i := range a {
		for _, c := range []int{2*i + 1, 2*i + 2} {
			if c < len(a) && a[c] < a[i] {
				return false
			}
		}
	}
	return true
}

// siftDown «топит» элемент i вниз, пока он больше меньшего из детей.
func siftDown(a []int, i int) {
	n := len(a)
	for {
		l, r, smallest := 2*i+1, 2*i+2, i
		if l < n && a[l] < a[smallest] {
			smallest = l
		}
		if r < n && a[r] < a[smallest] {
			smallest = r
		}
		if smallest == i {
			return
		}
		a[i], a[smallest] = a[smallest], a[i]
		i = smallest
	}
}

// heapify превращает срез в min-кучу на месте.
func heapify(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		siftDown(a, i)
	}
}

func main() {
	fmt.Println("--- Проверка 07_heapify_array (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := []int{9, 7, 5, 3, 1, 8, 6, 4, 2}
	heapify(a)
	check(isMinHeap(a), fmt.Sprintf("после heapify срез должен быть min-кучей, получено %v", a))
	check(len(a) > 0 && a[0] == 1, "корень min-кучи должен быть минимальным элементом (1)")

	sorted := append([]int(nil), a...)
	sort.Ints(sorted)
	check(fmt.Sprint(sorted) == "[1 2 3 4 5 6 7 8 9]", "heapify не должен терять или дублировать элементы")

	b := []int{3, 2}
	heapify(b)
	check(fmt.Sprint(b) == "[2 3]", fmt.Sprintf("два элемента: ожидалось [2 3], получено %v", b))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
