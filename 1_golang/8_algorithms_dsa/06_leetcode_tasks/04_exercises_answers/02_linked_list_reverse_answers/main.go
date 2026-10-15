// Упражнение 02: разворот списка: итеративно и рекурсивно (LeetCode 206) (Решение).
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
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// reverseRec разворачивает список рекурсивно.
func reverseRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	fmt.Println("--- Проверка 02_linked_list_reverse (Решение) ---")

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
