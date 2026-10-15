// Упражнение 08: обход дерева по уровням (BFS).
//
// Задача: реализуйте levelOrder(root) — значения дерева по уровням:
// результат [][]int, где i-й срез — узлы i-го уровня слева направо.
// Подсказка: очередь узлов; на каждой итерации обрабатывайте ровно
// len(queue) узлов — это и есть текущий уровень.
//
// Запуск: go run ./03_exercises/08_tree_level_order/main.go
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

// levelOrder возвращает значения дерева по уровням.
func levelOrder(root *TreeNode) [][]int {
	// TODO: BFS с разбиением на уровни.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_tree_level_order ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := levelOrder(buildBST(4, 2, 6, 1, 3, 5, 7))
	check(fmt.Sprint(got) == "[[4] [2 6] [1 3 5 7]]", fmt.Sprintf("ожидалось [[4] [2 6] [1 3 5 7]], получено %v", got))

	got = levelOrder(buildBST(1, 2, 3))
	check(fmt.Sprint(got) == "[[1] [2] [3]]", fmt.Sprintf("цепочка: ожидалось [[1] [2] [3]], получено %v", got))

	check(len(levelOrder(nil)) == 0, "пустое дерево: ожидался пустой результат")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
