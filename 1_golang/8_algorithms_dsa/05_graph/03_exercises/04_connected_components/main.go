// Упражнение 04: подсчёт компонент связности.
//
// Задача: реализуйте countComponents(n, edges) — число компонент связности
// неориентированного графа с вершинами 0..n-1.
// Варианты: BFS/DFS из каждой непосещённой вершины ИЛИ Union-Find (DSU).
// Изолированная вершина — отдельная компонента.
//
// Запуск: go run ./03_exercises/04_connected_components/main.go
package main

import "fmt"

// countComponents возвращает число компонент связности.
func countComponents(n int, edges [][2]int) int {
	// TODO: DFS/BFS или Union-Find.
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_connected_components ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := countComponents(5, [][2]int{{0, 1}, {1, 2}, {3, 4}})
	check(got == 2, fmt.Sprintf("{0,1,2} и {3,4}: ожидалось 2, получено %d", got))

	got = countComponents(4, nil)
	check(got == 4, fmt.Sprintf("4 вершины без рёбер: ожидалось 4, получено %d", got))

	got = countComponents(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	check(got == 1, fmt.Sprintf("треугольник: ожидалось 1, получено %d", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
