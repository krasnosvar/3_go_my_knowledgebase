// Упражнение 04: удаление узла по значению.
//
// Задача: реализуйте deleteValue(head, val) — удаляет ПЕРВЫЙ узел со значением
// val и возвращает (возможно новую) голову списка.
//   - удаление головы: новая голова — head.Next;
//   - удаление из середины/хвоста: prev.Next = cur.Next;
//   - если значения нет — список не меняется.
//
// Подсказка: фиктивный узел (dummy) перед головой убирает частный случай головы.
//
// Запуск: go run ./03_exercises/04_delete_node/main.go
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

// deleteValue удаляет первый узел со значением val и возвращает голову.
func deleteValue(head *Node, val int) *Node {
	// TODO: найдите узел и «перешейте» указатель предыдущего узла.
	return head
}

func main() {
	fmt.Println("--- Проверка 04_delete_node ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := toSlice(deleteValue(fromSlice([]int{1, 2, 3}), 1))
	check(fmt.Sprint(got) == "[2 3]", fmt.Sprintf("удаление головы: ожидалось [2 3], получено %v", got))

	got = toSlice(deleteValue(fromSlice([]int{1, 2, 3}), 2))
	check(fmt.Sprint(got) == "[1 3]", fmt.Sprintf("удаление из середины: ожидалось [1 3], получено %v", got))

	got = toSlice(deleteValue(fromSlice([]int{1, 2, 3}), 3))
	check(fmt.Sprint(got) == "[1 2]", fmt.Sprintf("удаление хвоста: ожидалось [1 2], получено %v", got))

	got = toSlice(deleteValue(fromSlice([]int{1, 2}), 9))
	check(fmt.Sprint(got) == "[1 2]", fmt.Sprintf("значения нет: ожидалось [1 2], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
