// Упражнение 01: добавление в конец списка.
//
// Задача: реализуйте appendNode(head, val) — добавляет узел со значением val
// в КОНЕЦ односвязного списка и возвращает голову списка.
//   - если список пуст (head == nil), новый узел становится головой;
//   - иначе дойдите до последнего узла и привяжите новый к его Next.
//
// Сложность: O(n) по времени, O(1) по памяти.
//
// Запуск: go run ./03_exercises/01_append/main.go
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

// appendNode добавляет val в конец списка и возвращает голову.
func appendNode(head *Node, val int) *Node {
	// TODO: обработайте пустой список, затем дойдите до хвоста.
	return head
}

func main() {
	fmt.Println("--- Проверка 01_append ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := toSlice(appendNode(nil, 1))
	check(fmt.Sprint(got) == "[1]", fmt.Sprintf("appendNode(nil, 1): ожидалось [1], получено %v", got))

	got = toSlice(appendNode(fromSlice([]int{1, 2}), 3))
	check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("appendNode([1 2], 3): ожидалось [1 2 3], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
