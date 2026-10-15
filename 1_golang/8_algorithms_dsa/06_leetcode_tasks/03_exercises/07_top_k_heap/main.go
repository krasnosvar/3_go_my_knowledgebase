// Упражнение 07: k самых частых элементов (LeetCode 347).
//
// Задача: реализуйте topKFrequent(nums, k) — k самых частых элементов.
// Порядок результата: по убыванию частоты, при равной частоте — по
// возрастанию значения (для детерминированной проверки).
// Подход: map частот + min-куча размера k по частоте (container/heap),
// O(n log k). Тип freqHeap (heap.Interface) уже готов.
//
// Запуск: go run ./03_exercises/07_top_k_heap/main.go
package main

import (
	"container/heap"
	"fmt"
)

// Держим импорт container/heap живым в заготовке.
var _ = heap.Init

type pair struct {
	val, count int
}

// freqHeap — min-куча: на вершине «наименее приоритетный» элемент
// (меньшая частота, при равенстве — большее значение).
type freqHeap []pair

func (h freqHeap) Len() int { return len(h) }
func (h freqHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	}
	return h[i].val > h[j].val
}
func (h freqHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Push(x any)   { *h = append(*h, x.(pair)) }
func (h *freqHeap) Pop() any {
	old := *h
	p := old[len(old)-1]
	*h = old[:len(old)-1]
	return p
}

// topKFrequent возвращает k самых частых элементов.
func topKFrequent(nums []int, k int) []int {
	// TODO: частоты -> куча размера k -> извлечение в обратном порядке.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_top_k_heap ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	check(fmt.Sprint(got) == "[1 2]", fmt.Sprintf("ожидалось [1 2], получено %v", got))

	got = topKFrequent([]int{4, 4, 5, 5, 6, 7, 7, 7}, 2)
	check(fmt.Sprint(got) == "[7 4]", fmt.Sprintf("равные частоты 4 и 5: ожидалось [7 4], получено %v", got))

	got = topKFrequent([]int{9}, 1)
	check(fmt.Sprint(got) == "[9]", fmt.Sprintf("ожидалось [9], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
