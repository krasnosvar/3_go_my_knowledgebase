// Упражнение 08: проверка двудольности (раскраска в 2 цвета).
//
// Задача: реализуйте isBipartite(adj) — можно ли раскрасить вершины
// неориентированного графа в 2 цвета так, чтобы концы каждого ребра были
// разного цвета. BFS/DFS: красим соседа в противоположный цвет; если сосед
// уже покрашен в тот же цвет — граф не двудольный.
// Граф может быть несвязным — запускайте обход из каждой некрашеной вершины.
// Факт: граф двудолен тогда и только тогда, когда в нём нет нечётных циклов.
//
// Запуск: go run ./03_exercises/08_is_bipartite/main.go
package main

import "fmt"

// isBipartite сообщает, двудолен ли граф.
func isBipartite(adj [][]int) bool {
	// TODO: раскраска 0/1, -1 — не покрашена.
	return false
}

func main() {
	fmt.Println("--- Проверка 08_is_bipartite ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	square := [][]int{0: {1, 3}, 1: {0, 2}, 2: {1, 3}, 3: {0, 2}}
	check(isBipartite(square), "цикл из 4 вершин двудолен")

	triangle := [][]int{0: {1, 2}, 1: {0, 2}, 2: {0, 1}}
	check(!isBipartite(triangle), "треугольник (нечётный цикл) не двудолен")

	mixed := [][]int{0: {1}, 1: {0}, 2: {3, 4}, 3: {2, 4}, 4: {2, 3}}
	check(!isBipartite(mixed), "несвязный граф с треугольником во второй компоненте не двудолен")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
