// Упражнение 01: построение списка смежности.
//
// Задача: реализуйте buildUndirected(n, edges) — список смежности [][]int для
// НЕОРИЕНТИРОВАННОГО графа с вершинами 0..n-1. Каждое ребро {u, v} добавляет
// v в adj[u] и u в adj[v]. Порядок соседей — порядок рёбер во входе.
// Вершины без рёбер должны иметь пустой (не nil) список соседей.
//
// Запуск: go run ./03_exercises/01_build_graph/main.go
package main

import "fmt"

// buildUndirected строит список смежности неориентированного графа.
func buildUndirected(n int, edges [][2]int) [][]int {
	// TODO: make([][]int, n) и добавление обоих направлений ребра.
	return nil
}

func main() {
	fmt.Println("--- Проверка 01_build_graph ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	adj := buildUndirected(5, [][2]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}})
	check(fmt.Sprint(adj) == "[[1 2] [0 2] [0 1 3] [2] []]",
		fmt.Sprintf("ожидалось [[1 2] [0 2] [0 1 3] [2] []], получено %v", adj))
	check(len(adj) == 5 && adj[4] != nil, "изолированная вершина 4 должна иметь пустой, но не nil список")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
