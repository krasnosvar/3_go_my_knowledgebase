// Упражнение 01: разбор аргументов через flag.FlagSet (Решение).
package main

import (
	"flag"
	"fmt"
	"io"
)

type options struct {
	N       int
	Verbose bool
	Files   []string
}

func parseArgs(args []string) (options, error) {
	var o options
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.IntVar(&o.N, "n", 10, "сколько строк")
	fs.BoolVar(&o.Verbose, "v", false, "подробный вывод")
	if err := fs.Parse(args); err != nil {
		return options{}, err
	}
	o.Files = fs.Args()
	return o, nil
}

func main() {
	fmt.Println("--- Проверка 01_arg_parse (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	o, err := parseArgs(nil)
	check(err == nil && o.N == 10 && !o.Verbose && len(o.Files) == 0, "без аргументов → дефолты")
	o, err = parseArgs([]string{"-n", "3", "-v", "a.txt", "b.txt"})
	check(err == nil && o.N == 3 && o.Verbose, "-n 3 -v разобраны")
	check(fmt.Sprint(o.Files) == "[a.txt b.txt]", "позиционные аргументы → Files")
	_, err = parseArgs([]string{"-x"})
	check(err != nil, "неизвестный флаг → ошибка")
	_, err = parseArgs([]string{"-n", "many"})
	check(err != nil, "-n many → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
