// Упражнение 05: разворот списка.
//
// Задача: реализуйте reverse(head) — разворачивает список на месте
// (без создания новых узлов) и возвращает новую голову.
// Классика с тремя указателями: prev, cur, next. O(n) времени, O(1) памяти.
//
// Запуск: go run ./03_exercises/05_reverse_list/main.go
package main

import "fmt"

// Node — узел односвязного списка.
type Node struct {
	Val  int
	Next *Node
}

// fromSlice строит список из среза (вспомогательная функция для теста).
func fromSlice(vals []int) *Node {
	var head *Node
	for i := len(vals) - 1; i >= 0; i-- {
		head = &Node{Val: vals[i], Next: head}
	}
	return head
}

// toSlice переводит список в срез; ограничение длины защищает тест от зацикливания.
func toSlice(head *Node) []int {
	var out []int
	for cur := head; cur != nil && len(out) < 1000; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

// reverse разворачивает список и возвращает новую голову.
func reverse(head *Node) *Node {
	// TODO: на каждом шаге разворачивайте cur.Next в сторону prev.
	return head
}

func main() {
	fmt.Println("--- Проверка 05_reverse_list ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := toSlice(reverse(fromSlice([]int{1, 2, 3, 4})))
	check(fmt.Sprint(got) == "[4 3 2 1]", fmt.Sprintf("reverse([1 2 3 4]): ожидалось [4 3 2 1], получено %v", got))

	got = toSlice(reverse(fromSlice([]int{7})))
	check(fmt.Sprint(got) == "[7]", fmt.Sprintf("reverse([7]): ожидалось [7], получено %v", got))

	check(reverse(nil) == nil, "reverse(nil) должен вернуть nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
