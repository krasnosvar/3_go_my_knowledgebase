// Упражнение 05: stdout для данных, stderr для диагностики (Решение).
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func upper(lines []string, stdout, stderr io.Writer) int {
	warns := 0
	for i, l := range lines {
		if l == "" {
			fmt.Fprintf(stderr, "warn: line %d is empty\n", i+1)
			warns++
			continue
		}
		fmt.Fprintln(stdout, strings.ToUpper(l))
	}
	return warns
}

func main() {
	fmt.Println("--- Проверка 05_stderr_output (Решение) ---")

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
