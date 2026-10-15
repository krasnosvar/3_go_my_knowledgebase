// Упражнение 03: PID-файл против двойного запуска.
//
// Задача:
//   - acquire(path string, pid int) error: атомарно создать файл
//     (os.OpenFile с флагами os.O_CREATE|os.O_EXCL|os.O_WRONLY, права 0o644)
//     и записать в него pid; если файл уже есть → errAlreadyRunning
//     (errors.Is(err, fs.ErrExist));
//   - release(path string) error: удалить файл.
//
// Проверка «существует ли» и создание должны быть ОДНОЙ операцией (O_EXCL),
// иначе два процесса могут стартовать одновременно.
//
// Запуск: go run ./03_exercises/03_pid_file
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

var _ = fs.ErrExist
var _ = strconv.Itoa

func acquire(path string, pid int) error {
	// TODO: os.OpenFile с O_EXCL, запись pid, Close.
	return nil
}

func release(path string) error {
	// TODO: os.Remove.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_pid_file ---")

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
