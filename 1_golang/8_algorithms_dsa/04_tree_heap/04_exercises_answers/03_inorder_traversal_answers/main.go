// Упражнение 03: симметричный (in-order) обход (Решение).
package main

import "fmt"

// TreeNode — узел двоичного дерева.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// insertBST вставляет значение в двоичное дерево поиска (вспомогательная функция для теста).
func insertBST(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if v < root.Val {
		root.Left = insertBST(root.Left, v)
	} else {
		root.Right = insertBST(root.Right, v)
	}
	return root
}

// buildBST строит BST, вставляя значения по порядку.
func buildBST(vals ...int) *TreeNode {
	var root *TreeNode
	for _, v := range vals {
		root = insertBST(root, v)
	}
	return root
}

// inorder возвращает значения дерева в порядке in-order.
func inorder(root *TreeNode) []int {
	var out []int
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out = append(out, cur.Val)
		cur = cur.Right
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 03_inorder_traversal (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := inorder(buildBST(5, 3, 8, 1, 4, 9, 7))
	check(fmt.Sprint(got) == "[1 3 4 5 7 8 9]", fmt.Sprintf("ожидалось [1 3 4 5 7 8 9], получено %v", got))

	got = inorder(buildBST(2))
	check(fmt.Sprint(got) == "[2]", fmt.Sprintf("один узел: ожидалось [2], получено %v", got))

	check(len(inorder(nil)) == 0, "пустое дерево: ожидался пустой результат")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
