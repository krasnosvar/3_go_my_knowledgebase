// Упражнение 04: временный файл и функция очистки.
//
// Задача: реализуйте newTempFile(content), которая:
//   - создаёт временный файл через os.CreateTemp("", "kb_*.txt");
//   - записывает в него content и закрывает файл;
//   - возвращает путь и функцию cleanup, удаляющую файл (os.Remove);
//   - при ошибке записи удаляет созданный файл сама и возвращает ошибку.
//
// Паттерн «ресурс + cleanup-функция» часто встречается в тестах
// (ср. t.Cleanup в пакете testing).
//
// Запуск: go run ./03_exercises/04_temp_file/main.go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// newTempFile создаёт временный файл с содержимым и возвращает cleanup.
func newTempFile(content string) (string, func(), error) {
	// TODO: os.CreateTemp, WriteString, Close, вернуть f.Name() и cleanup.
	return "", func() {}, nil
}

func main() {
	fmt.Println("--- Проверка 04_temp_file ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	path, cleanup, err := newTempFile("temporary data")
	if err != nil || cleanup == nil {
		fmt.Println("FAIL: newTempFile вернул ошибку или nil cleanup:", err)
		fmt.Println("FAIL")
		return
	}
	data, rerr := os.ReadFile(path)
	check(rerr == nil && string(data) == "temporary data", fmt.Sprintf("содержимое %q, err=%v", data, rerr))
	check(strings.HasPrefix(path, os.TempDir()), fmt.Sprintf("файл должен лежать в os.TempDir(), путь %q", path))

	cleanup()
	_, serr := os.Stat(path)
	check(path != "" && errors.Is(serr, fs.ErrNotExist), "после cleanup файл должен быть удалён")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
