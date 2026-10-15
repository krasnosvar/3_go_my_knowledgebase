// Упражнение 08: проверка существования файла (Решение).
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// fileExists сообщает, существует ли обычный файл по пути path.
func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("stat %s: %w", path, err)
	}
	return info.Mode().IsRegular(), nil
}

func main() {
	fmt.Println("--- Проверка 08_file_exists (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dir, err := os.MkdirTemp("", "file_exists_*")
	if err != nil {
		fmt.Println("FAIL: не удалось создать временную директорию:", err)
		return
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		fmt.Println("FAIL: подготовка файла:", err)
		return
	}

	ok, err := fileExists(file)
	check(ok && err == nil, fmt.Sprintf("существующий файл: получено %t, %v", ok, err))

	ok, err = fileExists(filepath.Join(dir, "missing.txt"))
	check(!ok && err == nil, fmt.Sprintf("несуществующий файл: ожидалось false, nil; получено %t, %v", ok, err))

	ok, err = fileExists(dir)
	check(!ok && err == nil, fmt.Sprintf("директория: ожидалось false, nil; получено %t, %v", ok, err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
