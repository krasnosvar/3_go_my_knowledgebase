// Упражнение 06: слияние двух отсортированных списков (Решение).
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

// mergeSorted сливает два отсортированных списка в один.
func mergeSorted(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

func main() {
	fmt.Println("--- Проверка 06_merge_sorted_lists (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := toSlice(mergeSorted(fromSlice([]int{1, 3, 5}), fromSlice([]int{2, 4, 6, 8})))
	check(fmt.Sprint(got) == "[1 2 3 4 5 6 8]", fmt.Sprintf("ожидалось [1 2 3 4 5 6 8], получено %v", got))

	got = toSlice(mergeSorted(nil, fromSlice([]int{1, 2})))
	check(fmt.Sprint(got) == "[1 2]", fmt.Sprintf("слияние с пустым: ожидалось [1 2], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
