// Упражнение 03: PID-файл против двойного запуска (Решение).
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

var errAlreadyRunning = errors.New("already running")

func acquire(path string, pid int) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if errors.Is(err, fs.ErrExist) {
		return errAlreadyRunning
	}
	if err != nil {
		return err
	}
	if _, err := f.WriteString(strconv.Itoa(pid)); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func release(path string) error {
	return os.Remove(path)
}

func main() {
	fmt.Println("--- Проверка 03_pid_file (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dir, err := os.MkdirTemp("", "pidfile")
	if err != nil {
		fmt.Println("FAIL: не удалось создать временный каталог:", err)
		return
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "app.pid")

	check(acquire(path, 4242) == nil, "первый запуск захватывает PID-файл")
	data, _ := os.ReadFile(path)
	check(string(data) == "4242", "в файле записан pid")
	check(errors.Is(acquire(path, 5555), errAlreadyRunning), "второй запуск → errAlreadyRunning")
	check(release(path) == nil, "release удаляет файл")
	_, statErr := os.Stat(path)
	check(errors.Is(statErr, fs.ErrNotExist), "файла больше нет")
	check(acquire(path, 6) == nil, "после release можно стартовать снова")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
