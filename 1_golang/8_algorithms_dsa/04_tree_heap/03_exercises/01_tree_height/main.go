// Упражнение 01: высота двоичного дерева.
//
// Задача: реализуйте height(root) — число уровней дерева (длина самого
// длинного пути от корня до листа в узлах). Пустое дерево — 0, один узел — 1.
// Рекурсия: height = 1 + max(height(left), height(right)).
//
// Запуск: go run ./03_exercises/01_tree_height/main.go
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

// height возвращает высоту дерева в узлах.
func height(root *TreeNode) int {
	// TODO: рекурсивно посчитайте высоту поддеревьев.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_tree_height ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(height(nil) == 0, "пустое дерево: ожидалась высота 0")
	check(height(buildBST(5)) == 1, fmt.Sprintf("один узел: ожидалась высота 1, получено %d", height(buildBST(5))))

	balanced := buildBST(4, 2, 6, 1, 3, 5, 7)
	check(height(balanced) == 3, fmt.Sprintf("сбалансированное дерево из 7 узлов: ожидалась высота 3, получено %d", height(balanced)))

	chain := buildBST(1, 2, 3, 4)
	check(height(chain) == 4, fmt.Sprintf("вырожденное дерево-цепочка: ожидалась высота 4, получено %d", height(chain)))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
