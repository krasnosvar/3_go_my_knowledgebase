// Упражнение 03: запись файла через bufio.Writer (Решение).
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// writeLines записывает строки в файл path.
func writeLines(path string, lines []string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close %s: %w", path, cerr)
		}
	}()
	w := bufio.NewWriter(f)
	for _, line := range lines {
		if _, err := w.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return w.Flush()
}

func main() {
	fmt.Println("--- Проверка 03_write_file (Решение) ---")

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
