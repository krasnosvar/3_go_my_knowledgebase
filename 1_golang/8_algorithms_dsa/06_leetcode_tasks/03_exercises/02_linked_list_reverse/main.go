// Упражнение 02: разворот списка: итеративно и рекурсивно (LeetCode 206).
//
// Задача: реализуйте два варианта разворота односвязного списка:
//   - reverseIter — итеративно, три указателя, O(1) памяти;
//   - reverseRec — рекурсивно: развернуть хвост, затем head.Next.Next = head,
//     head.Next = nil. Память O(n) из-за стека вызовов — частый вопрос интервью.
//
// Запуск: go run ./03_exercises/02_linked_list_reverse/main.go
package main

import "fmt"

// ListNode — узел односвязного списка (как в LeetCode).
type ListNode struct {
	Val  int
	Next *ListNode
}

func build(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func values(head *ListNode) []int {
	var out []int
	for cur := head; cur != nil && len(out) < 1000; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

// reverseIter разворачивает список итеративно.
func reverseIter(head *ListNode) *ListNode {
	// TODO: prev/cur/next.
	return head
}

// reverseRec разворачивает список рекурсивно.
func reverseRec(head *ListNode) *ListNode {
	// TODO: база рекурсии — пустой список или один узел.
	return head
}

func main() {
	fmt.Println("--- Проверка 02_linked_list_reverse ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := values(reverseIter(build(1, 2, 3, 4, 5)))
	check(fmt.Sprint(got) == "[5 4 3 2 1]", fmt.Sprintf("reverseIter: ожидалось [5 4 3 2 1], получено %v", got))

	got = values(reverseRec(build(1, 2, 3, 4, 5)))
	check(fmt.Sprint(got) == "[5 4 3 2 1]", fmt.Sprintf("reverseRec: ожидалось [5 4 3 2 1], получено %v", got))

	check(reverseIter(nil) == nil && reverseRec(nil) == nil, "разворот пустого списка должен вернуть nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
