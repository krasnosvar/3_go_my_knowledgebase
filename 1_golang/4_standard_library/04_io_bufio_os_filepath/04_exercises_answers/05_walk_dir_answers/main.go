// Упражнение 05: обход дерева директорий (Решение).
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// findByExt возвращает относительные пути файлов с расширением ext.
func findByExt(root, ext string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ext {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		out = append(out, rel)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s: %w", root, err)
	}
	return out, nil
}

func main() {
	fmt.Println("--- Проверка 05_walk_dir (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	root, err := os.MkdirTemp("", "walk_dir_*")
	if err != nil {
		fmt.Println("FAIL: не удалось создать временную директорию:", err)
		return
	}
	defer os.RemoveAll(root)

	for _, p := range []string{"main.go", "README.md", "pkg/a.go", "pkg/b.txt", "pkg/sub/c.go", "dir.go/x.md"} {
		full := filepath.Join(root, filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			fmt.Println("FAIL: подготовка дерева:", err)
			return
		}
		if err := os.WriteFile(full, []byte("x"), 0o644); err != nil {
			fmt.Println("FAIL: подготовка дерева:", err)
			return
		}
	}

	got, err := findByExt(root, ".go")
	want := []string{"main.go", filepath.Join("pkg", "a.go"), filepath.Join("pkg", "sub", "c.go")}
	check(err == nil && fmt.Sprint(got) == fmt.Sprint(want), fmt.Sprintf("ожидалось %v, получено %v (err=%v)", want, got, err))

	_, err = findByExt(filepath.Join(root, "missing"), ".go")
	check(err != nil, "несуществующий root должен давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
