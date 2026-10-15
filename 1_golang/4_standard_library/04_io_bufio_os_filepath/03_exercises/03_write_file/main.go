// Упражнение 03: запись файла через bufio.Writer.
//
// Задача: реализуйте writeLines(path, lines), которая создаёт (или
// перезаписывает) файл и пишет в него строки, каждую с '\n' на конце.
//   - os.Create + defer f.Close();
//   - запись через bufio.NewWriter(f) — и ОБЯЗАТЕЛЬНО w.Flush() в конце,
//     иначе хвост данных останется в буфере;
//   - ошибку f.Close() тоже стоит вернуть: для записи она значима.
//
// Тест работает во временной директории и удаляет её за собой.
//
// Запуск: go run ./03_exercises/03_write_file/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var _ = bufio.NewWriter // держим импорт живым в заготовке

// writeLines записывает строки в файл path.
func writeLines(path string, lines []string) (err error) {
	// TODO: os.Create, bufio.NewWriter, WriteString(line + "\n"), Flush, Close.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_write_file ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dir, err := os.MkdirTemp("", "write_file_*")
	if err != nil {
		fmt.Println("FAIL: не удалось создать временную директорию:", err)
		return
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "out.txt")
	err = writeLines(path, []string{"alpha", "beta", "gamma"})
	data, rerr := os.ReadFile(path)
	check(err == nil && rerr == nil && string(data) == "alpha\nbeta\ngamma\n",
		fmt.Sprintf("содержимое %q, err=%v, readErr=%v", data, err, rerr))

	err = writeLines(path, []string{"new"})
	data, _ = os.ReadFile(path)
	check(err == nil && string(data) == "new\n", fmt.Sprintf("повторная запись должна перезаписать файл, получено %q", data))

	err = writeLines(filepath.Join(dir, "no_such_dir", "x.txt"), []string{"x"})
	check(err != nil, "запись в несуществующую директорию должна давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
