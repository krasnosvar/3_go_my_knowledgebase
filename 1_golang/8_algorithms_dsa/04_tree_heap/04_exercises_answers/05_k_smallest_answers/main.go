// Упражнение 05: k наименьших элементов через кучу (Решение).
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// maxHeap — max-куча для container/heap.
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// kSmallest возвращает k наименьших элементов по возрастанию.
func kSmallest(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	h := &maxHeap{}
	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v < (*h)[0] {
			(*h)[0] = v
			heap.Fix(h, 0)
		}
	}
	out := append([]int(nil), (*h)...)
	sort.Ints(out)
	return out
}

func main() {
	fmt.Println("--- Проверка 05_k_smallest (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := kSmallest([]int{7, 2, 9, 4, 1, 8, 3}, 3)
	check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("k=3: ожидалось [1 2 3], получено %v", got))

	got = kSmallest([]int{5, 4}, 5)
	check(fmt.Sprint(got) == "[4 5]", fmt.Sprintf("k > len: ожидалось [4 5], получено %v", got))

	got = kSmallest([]int{1, 2}, 0)
	check(got != nil && len(got) == 0, fmt.Sprintf("k=0: ожидался пустой срез, получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
