// Упражнение 08: число островов на сетке (LeetCode 200).
//
// Задача: реализуйте numIslands(grid) — число островов: групп '1',
// соединённых по горизонтали/вертикали. '0' — вода.
// Подход: проходим по клеткам; на каждой непосещённой '1' увеличиваем счётчик
// и запускаем BFS (очередь координат), помечая весь остров посещённым.
// Не меняйте входную сетку — используйте отдельный visited.
//
// Запуск: go run ./03_exercises/08_graph_bfs/main.go
package main

import "fmt"

// numIslands возвращает число островов на сетке.
func numIslands(grid [][]byte) int {
	// TODO: перебор клеток + BFS по 4 направлениям.
	return 0
}

func main() {
	fmt.Println("--- Проверка 08_graph_bfs ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	got := numIslands(grid)
	check(got == 3, fmt.Sprintf("ожидалось 3 острова, получено %d", got))
	check(string(grid[0]) == "11000", "входная сетка не должна изменяться")

	diag := [][]byte{[]byte("101"), []byte("010"), []byte("101")}
	got = numIslands(diag)
	check(got == 5, fmt.Sprintf("диагональ не соединяет клетки: ожидалось 5, получено %d", got))

	check(numIslands(nil) == 0, "пустая сетка: ожидалось 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
