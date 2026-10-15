// Упражнение 08: help и -h (Решение).
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(stdout)
	fs.String("out", "", "куда писать результат")
	fs.Usage = func() {
		fmt.Fprintln(stdout, "Usage: tool [flags] <input>")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() == 0 {
		fs.Usage()
		return 2
	}
	return 0
}

func main() {
	fmt.Println("--- Проверка 08_help_text (Решение) ---")

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
