// Упражнение 05: stdout для данных, stderr для диагностики.
//
// Задача: реализуйте upper(lines []string, stdout, stderr io.Writer) int:
//   - каждую непустую строку печатать в stdout в верхнем регистре + "\n";
//   - на пустую строку печатать в stderr "warn: line N is empty\n"
//     (N — номер с 1) и продолжать;
//   - вернуть число предупреждений.
//
// Так `tool | next-tool` получает только данные, а предупреждения видит человек.
//
// Запуск: go run ./03_exercises/05_stderr_output
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var _ = strings.ToUpper

func upper(lines []string, stdout, stderr io.Writer) int {
	// TODO: данные → stdout, предупреждения → stderr.
	return 0
}

func main() {
	fmt.Println("--- Проверка 05_stderr_output ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var out, errOut bytes.Buffer
	n := upper([]string{"a", "", "bc"}, &out, &errOut)
	check(out.String() == "A\nBC\n", "stdout содержит только данные")
	check(errOut.String() == "warn: line 2 is empty\n", "stderr содержит предупреждение")
	check(n == 1, "одно предупреждение")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
