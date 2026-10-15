// Упражнение 03: поиск узла по значению (Решение).
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

// find возвращает первый узел со значением val или nil.
func find(head *Node, val int) *Node {
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			return cur
		}
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_find_node (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	head := fromSlice([]int{5, 7, 9})
	n := find(head, 7)
	check(n != nil && n.Val == 7 && n.Next != nil && n.Next.Val == 9, "find(7) должен вернуть узел 7, за которым идёт 9")
	check(find(head, 42) == nil, "find(42) должен вернуть nil")
	check(find(nil, 1) == nil, "find в пустом списке должен вернуть nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
