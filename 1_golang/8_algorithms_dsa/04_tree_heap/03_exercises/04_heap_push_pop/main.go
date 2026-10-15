// Упражнение 04: min-heap на срезе вручную.
//
// Задача: реализуйте двоичную min-кучу на срезе БЕЗ container/heap:
//   - Push(v): добавить в конец и «всплыть» (sift up): пока родитель
//     (i-1)/2 больше — меняем местами;
//   - Pop(): вернуть корень data[0], перенести последний элемент в корень
//     и «утопить» (sift down) к меньшему из детей 2i+1, 2i+2;
//   - пустая куча: Pop возвращает (0, false).
//
// Обе операции O(log n).
//
// Запуск: go run ./03_exercises/04_heap_push_pop/main.go
package main

import "fmt"

// MinHeap — двоичная min-куча на срезе.
type MinHeap struct {
	data []int
}

func (h *MinHeap) Push(v int) {
	// TODO: append + sift up.
}

func (h *MinHeap) Pop() (int, bool) {
	// TODO: забрать корень, последний в корень, sift down.
	return 0, false
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func main() {
	fmt.Println("--- Проверка 04_heap_push_pop ---")

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
