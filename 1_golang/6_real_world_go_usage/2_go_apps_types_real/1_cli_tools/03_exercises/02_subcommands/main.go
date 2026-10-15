// Упражнение 02: подкоманды без сторонних библиотек.
//
// Задача: реализуйте dispatch(args []string, out io.Writer) int:
//   - "version"          → печатает "tool v1.0.0\n", код 0;
//   - "greet -name X"    → печатает "hello, X\n" (у greet свой FlagSet,
//     -name по умолчанию "world"), код 0;
//   - пустой args или неизвестная команда → печатает "usage: tool <version|greet>\n", код 2;
//   - ошибка флагов подкоманды → код 2.
//
// Запуск: go run ./03_exercises/02_subcommands
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
)

var _ = flag.NewFlagSet

func dispatch(args []string, out io.Writer) int {
	// TODO: switch по args[0], у greet — свой FlagSet.
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_subcommands ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	run := func(args ...string) (string, int) {
		var b bytes.Buffer
		code := dispatch(args, &b)
		return b.String(), code
	}

	out, code := run("version")
	check(code == 0 && out == "tool v1.0.0\n", "version")
	out, code = run("greet", "-name", "den")
	check(code == 0 && out == "hello, den\n", "greet -name den")
	out, code = run("greet")
	check(code == 0 && out == "hello, world\n", "greet без флага")
	out, code = run("deploy")
	check(code == 2 && out == "usage: tool <version|greet>\n", "неизвестная команда → usage, 2")
	_, code = run()
	check(code == 2, "без команды → 2")
	_, code = run("greet", "-bad")
	check(code == 2, "ошибка флага подкоманды → 2")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
