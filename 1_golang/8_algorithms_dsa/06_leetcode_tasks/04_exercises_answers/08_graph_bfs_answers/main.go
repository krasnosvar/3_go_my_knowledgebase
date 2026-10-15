// Упражнение 08: число островов на сетке (LeetCode 200) (Решение).
package main

import "fmt"

// numIslands возвращает число островов на сетке.
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '1' || visited[r][c] {
				continue
			}
			count++
			visited[r][c] = true
			queue := [][2]int{{r, c}}
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for _, d := range dirs {
					nr, nc := cur[0]+d[0], cur[1]+d[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' && !visited[nr][nc] {
						visited[nr][nc] = true
						queue = append(queue, [2]int{nr, nc})
					}
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println("--- Проверка 08_graph_bfs (Решение) ---")

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
