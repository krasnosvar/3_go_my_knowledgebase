// Упражнение 02: построчное чтение через bufio.Scanner.
//
// Задача: реализуйте countNonEmpty(r), которая читает r построчно
// (bufio.NewScanner) и считает непустые строки:
//   - строка из одних пробелов/табов считается пустой (strings.TrimSpace);
//   - строки, начинающиеся с '#' (после TrimSpace), — комментарии, не считаются;
//   - после цикла обязательно проверьте scanner.Err().
//
// Запуск: go run ./03_exercises/02_scanner_lines/main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var _ = bufio.NewScanner // держим импорт живым в заготовке

// countNonEmpty считает непустые строки без комментариев.
func countNonEmpty(r io.Reader) (int, error) {
	// TODO: bufio.NewScanner(r), цикл sc.Scan(), проверка sc.Err().
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 02_scanner_lines ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := "# config\nhost=localhost\n\n   \nport=8080\n  # disabled=true\ndebug=false"
	n, err := countNonEmpty(strings.NewReader(in))
	check(err == nil && n == 3, fmt.Sprintf("ожидалось 3, получено %d (err=%v)", n, err))

	n, err = countNonEmpty(strings.NewReader(""))
	check(err == nil && n == 0, fmt.Sprintf("пустой ввод: получено %d, %v", n, err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
