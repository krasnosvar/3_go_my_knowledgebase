// Упражнение 08: help и -h.
//
// Задача: реализуйте run(args []string, stdout, stderr io.Writer) int:
//   - FlagSet "tool" (ContinueOnError), флаг -out string "куда писать результат";
//   - fs.Usage печатает в stdout строку "Usage: tool [flags] <input>\n"
//     и затем fs.PrintDefaults() (fs.SetOutput(stdout) для help);
//   - -h / -help → Parse вернёт flag.ErrHelp: это НЕ ошибка, код 0;
//   - другая ошибка флагов → код 2;
//   - нет позиционного <input> → вызвать fs.Usage(), код 2;
//   - иначе код 0.
//
// Запуск: go run ./03_exercises/08_help_text
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

var _ = errors.Is
var _ = flag.ErrHelp

func run(args []string, stdout, stderr io.Writer) int {
	// TODO: настройте FlagSet, Usage и обработайте flag.ErrHelp.
	return -1
}

func main() {
	fmt.Println("--- Проверка 08_help_text ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var out bytes.Buffer
	code := run([]string{"-h"}, &out, io.Discard)
	help := out.String()
	check(code == 0, "-h → код 0")
	check(strings.HasPrefix(help, "Usage: tool [flags] <input>\n"), "help начинается со строки Usage")
	check(strings.Contains(help, "-out") && strings.Contains(help, "куда писать результат"), "help содержит описание -out")

	out.Reset()
	check(run(nil, &out, io.Discard) == 2 && strings.Contains(out.String(), "Usage:"), "без input → usage и код 2")
	check(run([]string{"-zzz"}, io.Discard, io.Discard) == 2, "неизвестный флаг → 2")
	check(run([]string{"-out", "x", "in.txt"}, io.Discard, io.Discard) == 0, "нормальный вызов → 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
