// Упражнение 06: проверка размера каталога.
//
// Задача: реализуйте dirSize(root string) (int64, error) — сумма размеров
// всех обычных файлов в дереве (filepath.WalkDir, d.Info().Size()), и
// checkQuota(root string, limit int64) (bool, error) — true, если размер
// превышает limit (пора слать алерт / чистить).
//
// Запуск: go run ./03_exercises/06_disk_check
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var _ fs.DirEntry

func dirSize(root string) (int64, error) {
	// TODO: filepath.WalkDir, суммируйте размеры обычных файлов.
	return 0, nil
}

func checkQuota(root string, limit int64) (bool, error) {
	// TODO: используйте dirSize.
	return false, nil
}

func main() {
	fmt.Println("--- Проверка 06_disk_check ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dir, err := os.MkdirTemp("", "diskcheck")
	if err != nil {
		fmt.Println("FAIL: не удалось создать временный каталог:", err)
		return
	}
	defer os.RemoveAll(dir)
	_ = os.MkdirAll(filepath.Join(dir, "sub", "deep"), 0o755)
	_ = os.WriteFile(filepath.Join(dir, "a.log"), make([]byte, 100), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "sub", "b.log"), make([]byte, 250), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "sub", "deep", "c.log"), make([]byte, 50), 0o644)

	size, err := dirSize(dir)
	check(err == nil && size == 400, fmt.Sprintf("размер дерева 400 байт (получено %d)", size))
	over, _ := checkQuota(dir, 300)
	check(over, "400 > 300 → превышение")
	over, _ = checkQuota(dir, 400)
	check(!over, "400 при лимите 400 → не превышение")
	_, err = dirSize(filepath.Join(dir, "missing"))
	check(err != nil, "несуществующий каталог → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
