// Упражнение 02: поиск в двоичном дереве поиска (Решение).
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

// search сообщает, есть ли target в BST.
func search(root *TreeNode, target int) bool {
	for cur := root; cur != nil; {
		switch {
		case target == cur.Val:
			return true
		case target < cur.Val:
			cur = cur.Left
		default:
			cur = cur.Right
		}
	}
	return false
}

func main() {
	fmt.Println("--- Проверка 02_bst_search (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	root := buildBST(8, 3, 10, 1, 6, 14, 4, 7, 13)
	for _, v := range []int{8, 1, 7, 13} {
		check(search(root, v), fmt.Sprintf("search(%d): ожидалось true", v))
	}
	for _, v := range []int{0, 5, 15} {
		check(!search(root, v), fmt.Sprintf("search(%d): ожидалось false", v))
	}
	check(!search(nil, 1), "поиск в пустом дереве: ожидалось false")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
