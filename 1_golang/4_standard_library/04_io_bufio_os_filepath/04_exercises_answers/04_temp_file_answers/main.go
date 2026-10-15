// Упражнение 04: временный файл и функция очистки (Решение).
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
	f, err := os.CreateTemp("", "kb_*.txt")
	if err != nil {
		return "", nil, fmt.Errorf("create temp: %w", err)
	}
	path := f.Name()
	if _, err := f.WriteString(content); err != nil {
		f.Close()
		os.Remove(path)
		return "", nil, fmt.Errorf("write temp: %w", err)
	}
	if err := f.Close(); err != nil {
		os.Remove(path)
		return "", nil, fmt.Errorf("close temp: %w", err)
	}
	return path, func() { os.Remove(path) }, nil
}

func main() {
	fmt.Println("--- Проверка 04_temp_file (Решение) ---")

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
