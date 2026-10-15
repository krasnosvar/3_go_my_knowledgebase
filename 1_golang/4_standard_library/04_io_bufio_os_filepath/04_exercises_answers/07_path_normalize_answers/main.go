// Упражнение 07: безопасная нормализация пути (защита от path traversal) (Решение).
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// safeJoin соединяет base и userPath, запрещая выход за base.
func safeJoin(base, userPath string) (string, error) {
	full := filepath.Join(base, userPath)
	rel, err := filepath.Rel(base, full)
	if err != nil {
		return "", fmt.Errorf("rel %q: %w", userPath, err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("path %q escapes base %q", userPath, base)
	}
	return full, nil
}

func main() {
	fmt.Println("--- Проверка 07_path_normalize (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	base := filepath.Join("srv", "static")

	got, err := safeJoin(base, "css/./site.css")
	want := filepath.Join(base, "css", "site.css")
	check(err == nil && got == want, fmt.Sprintf("ожидалось %q, получено %q (err=%v)", want, got, err))

	got, err = safeJoin(base, "img/../index.html")
	want = filepath.Join(base, "index.html")
	check(err == nil && got == want, fmt.Sprintf("ожидалось %q, получено %q (err=%v)", want, got, err))

	for _, bad := range []string{"../secret.txt", "../../etc/passwd", "css/../../.."} {
		if p, err := safeJoin(base, bad); err == nil {
			check(false, fmt.Sprintf("%q выходит за base, но получено %q без ошибки", bad, p))
		}
	}

	got, err = safeJoin(base, "..data/file")
	check(err == nil && got == filepath.Join(base, "..data", "file"), fmt.Sprintf("имя \"..data\" — обычная директория, получено %q, %v", got, err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
