// Упражнение 02: подкоманды без сторонних библиотек (Решение).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
)

func dispatch(args []string, out io.Writer) int {
	usage := func() int {
		fmt.Fprintln(out, "usage: tool <version|greet>")
		return 2
	}
	if len(args) == 0 {
		return usage()
	}
	switch args[0] {
	case "version":
		fmt.Fprintln(out, "tool v1.0.0")
		return 0
	case "greet":
		fs := flag.NewFlagSet("greet", flag.ContinueOnError)
		fs.SetOutput(io.Discard)
		name := fs.String("name", "world", "кого приветствовать")
		if err := fs.Parse(args[1:]); err != nil {
			return 2
		}
		fmt.Fprintf(out, "hello, %s\n", *name)
		return 0
	default:
		return usage()
	}
}

func main() {
	fmt.Println("--- Проверка 02_subcommands (Решение) ---")

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
