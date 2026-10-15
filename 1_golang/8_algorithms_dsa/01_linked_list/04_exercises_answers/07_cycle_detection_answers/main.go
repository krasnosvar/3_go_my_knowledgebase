// Упражнение 07: обнаружение цикла (Floyd) (Решение).
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

// hasCycle сообщает, есть ли в списке цикл.
func hasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("--- Проверка 07_cycle_detection (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	noCycle := fromSlice([]int{1, 2, 3, 4})
	check(!hasCycle(noCycle), "список без цикла: ожидалось false")
	check(!hasCycle(nil), "пустой список: ожидалось false")

	withCycle := fromSlice([]int{1, 2, 3, 4})
	last := withCycle
	for last.Next != nil {
		last = last.Next
	}
	last.Next = withCycle.Next // 4 -> 2
	check(hasCycle(withCycle), "список с циклом 4->2: ожидалось true")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
