// Упражнение 08: поиск середины списка (Решение).
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

// middle возвращает средний узел списка.
func middle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	fmt.Println("--- Проверка 08_middle_node (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	m := middle(fromSlice([]int{1, 2, 3, 4, 5}))
	check(m != nil && m.Val == 3, "[1 2 3 4 5]: ожидался узел 3")

	m = middle(fromSlice([]int{1, 2, 3, 4}))
	check(m != nil && m.Val == 3, "[1 2 3 4]: ожидался второй средний узел 3")

	m = middle(fromSlice([]int{9}))
	check(m != nil && m.Val == 9, "[9]: ожидался узел 9")

	check(middle(nil) == nil, "пустой список: ожидался nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
