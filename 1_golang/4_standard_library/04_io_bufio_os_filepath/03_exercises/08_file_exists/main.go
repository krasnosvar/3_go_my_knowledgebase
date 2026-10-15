// Упражнение 08: проверка существования файла.
//
// Задача: реализуйте fileExists(path), которая возвращает:
//   - (true, nil) — путь существует и это обычный файл;
//   - (false, nil) — путь не существует (errors.Is(err, fs.ErrNotExist))
//     или это директория;
//   - (false, err) — любая другая ошибка os.Stat (например, нет прав).
//
// Типичная ошибка — считать «любая ошибка Stat = файла нет».
//
// Запуск: go run ./03_exercises/08_file_exists/main.go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var _ = errors.Is // держим импорты живыми в заготовке
var _ = fs.ErrNotExist

// fileExists сообщает, существует ли обычный файл по пути path.
func fileExists(path string) (bool, error) {
	// TODO: os.Stat, разбор ошибки через errors.Is(err, fs.ErrNotExist), info.Mode().IsRegular().
	return true, nil
}

func main() {
	fmt.Println("--- Проверка 08_file_exists ---")

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
