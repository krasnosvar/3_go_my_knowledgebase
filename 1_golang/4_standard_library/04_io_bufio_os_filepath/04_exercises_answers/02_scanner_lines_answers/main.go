// Упражнение 02: построчное чтение через bufio.Scanner (Решение).
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// countNonEmpty считает непустые строки без комментариев.
func countNonEmpty(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	n := 0
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		n++
	}
	if err := sc.Err(); err != nil {
		return 0, fmt.Errorf("scan: %w", err)
	}
	return n, nil
}

func main() {
	fmt.Println("--- Проверка 02_scanner_lines (Решение) ---")

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
