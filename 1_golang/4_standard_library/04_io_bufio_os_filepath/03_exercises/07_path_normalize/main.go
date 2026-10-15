// Упражнение 07: безопасная нормализация пути (защита от path traversal).
//
// Задача: реализуйте safeJoin(base, userPath), которая строит путь к файлу
// внутри base по пути из запроса пользователя:
//   - результат = filepath.Join(base, userPath) (Join сам делает Clean);
//   - если результат выходит за пределы base (например, "../../etc/passwd"),
//     верните ошибку. Проверка: filepath.Rel(base, result) не должен
//     начинаться с ".." ;
//   - сам base ("" или ".") допустим.
//
// Запуск: go run ./03_exercises/07_path_normalize/main.go
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

var _ = strings.HasPrefix // держим импорт живым в заготовке

// safeJoin соединяет base и userPath, запрещая выход за base.
func safeJoin(base, userPath string) (string, error) {
	// TODO: filepath.Join + filepath.Rel + проверка префикса "..".
	return filepath.Join(base, userPath), nil
}

func main() {
	fmt.Println("--- Проверка 07_path_normalize ---")

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
