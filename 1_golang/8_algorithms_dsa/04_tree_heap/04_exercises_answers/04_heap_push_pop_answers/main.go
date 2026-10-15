// Упражнение 04: min-heap на срезе вручную (Решение).
package main

import "fmt"

// MinHeap — двоичная min-куча на срезе.
type MinHeap struct {
	data []int
}

func (h *MinHeap) Push(v int) {
	h.data = append(h.data, v)
	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) Pop() (int, bool) {
	n := len(h.data)
	if n == 0 {
		return 0, false
	}
	top := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	i := 0
	for {
		l, r, smallest := 2*i+1, 2*i+2, i
		if l < len(h.data) && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < len(h.data) && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
	return top, true
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func main() {
	fmt.Println("--- Проверка 04_heap_push_pop (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var h MinHeap
	_, ok := h.Pop()
	check(!ok, "Pop на пустой куче должен вернуть ok=false")

	for _, v := range []int{5, 1, 8, 3, 2, 9, 1} {
		h.Push(v)
	}
	var got []int
	for h.Len() > 0 && len(got) < 20 {
		v, _ := h.Pop()
		got = append(got, v)
	}
	check(fmt.Sprint(got) == "[1 1 2 3 5 8 9]", fmt.Sprintf("ожидалось [1 1 2 3 5 8 9], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
