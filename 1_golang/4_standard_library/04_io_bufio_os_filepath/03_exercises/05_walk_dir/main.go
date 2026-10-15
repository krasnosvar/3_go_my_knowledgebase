// Упражнение 05: обход дерева директорий.
//
// Задача: реализуйте findByExt(root, ext), которая обходит root через
// filepath.WalkDir и возвращает пути файлов с расширением ext
// (filepath.Ext(path) == ext), ОТНОСИТЕЛЬНО root (filepath.Rel),
// в порядке обхода (WalkDir идёт в лексическом порядке).
//   - директории в результат не попадают;
//   - ошибку из колбэка (err != nil) возвращайте наружу.
//
// Тест строит дерево во временной директории и удаляет его за собой.
//
// Запуск: go run ./03_exercises/05_walk_dir/main.go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var _ fs.DirEntry // держим импорт живым в заготовке

// findByExt возвращает относительные пути файлов с расширением ext.
func findByExt(root, ext string) ([]string, error) {
	// TODO: filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error { ... }).
	return nil, nil
}

func main() {
	fmt.Println("--- Проверка 05_walk_dir ---")

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
