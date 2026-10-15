// Упражнение 01: разбор аргументов через flag.FlagSet.
//
// Задача: реализуйте parseArgs(args []string) (options, error):
//   - собственный flag.NewFlagSet("tool", flag.ContinueOnError) (а не глобальный flag),
//     вывод ошибок пакета flag направьте в io.Discard;
//   - флаги: -n int (по умолчанию 10), -v bool;
//   - позиционные аргументы после флагов → options.Files;
//   - неизвестный флаг или битое значение → ошибка (а не os.Exit).
//
// Свой FlagSet + ContinueOnError делает CLI тестируемым.
//
// Запуск: go run ./03_exercises/01_arg_parse
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

var _ = flag.ContinueOnError
var _ = io.Discard

func parseArgs(args []string) (options, error) {
	// TODO: создайте FlagSet, объявите -n и -v, вызовите Parse(args), заберите fs.Args().
	return options{}, nil
}

func main() {
	fmt.Println("--- Проверка 01_arg_parse ---")

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
