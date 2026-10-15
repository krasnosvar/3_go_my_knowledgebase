// Упражнение 02: добавление в начало списка.
//
// Задача: реализуйте prepend(head, val) — добавляет узел со значением val
// в НАЧАЛО списка и возвращает новую голову.
// Это главное преимущество списка перед срезом: вставка в начало за O(1).
//
// Запуск: go run ./03_exercises/02_prepend/main.go
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

// prepend добавляет val в начало списка и возвращает новую голову.
func prepend(head *Node, val int) *Node {
	// TODO: создайте узел, у которого Next указывает на старую голову.
	return head
}

func main() {
	fmt.Println("--- Проверка 02_prepend ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := toSlice(prepend(nil, 1))
	check(fmt.Sprint(got) == "[1]", fmt.Sprintf("prepend(nil, 1): ожидалось [1], получено %v", got))

	got = toSlice(prepend(fromSlice([]int{2, 3}), 1))
	check(fmt.Sprint(got) == "[1 2 3]", fmt.Sprintf("prepend([2 3], 1): ожидалось [1 2 3], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
