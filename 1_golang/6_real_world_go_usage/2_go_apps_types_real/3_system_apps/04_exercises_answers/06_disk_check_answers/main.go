// Упражнение 06: проверка размера каталога (Решение).
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func dirSize(root string) (int64, error) {
	var total int64
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.Type().IsRegular() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		total += info.Size()
		return nil
	})
	return total, err
}

func checkQuota(root string, limit int64) (bool, error) {
	size, err := dirSize(root)
	if err != nil {
		return false, err
	}
	return size > limit, nil
}

func main() {
	fmt.Println("--- Проверка 06_disk_check (Решение) ---")

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
