// Упражнение 06: слияние конфига из файла и флагов.
//
// Задача: реализуйте merge(file config, args []string) (config, error):
//   - стартуем со значений из файла (file);
//   - флаги -host и -port (свой FlagSet) перекрывают файл, только если
//     они ЯВНО переданы (fs.Visit) — иначе дефолт флага не должен
//     затирать значение из файла;
//   - ошибка разбора флагов → ошибка.
//
// Запуск: go run ./03_exercises/06_config_merge
package main

import (
	"flag"
	"fmt"
	"io"
)

type config struct {
	Host string
	Port int
}

var _ = flag.NewFlagSet
var _ = io.Discard

func merge(file config, args []string) (config, error) {
	// TODO: разберите флаги и примените только явно заданные.
	return file, nil
}

func main() {
	fmt.Println("--- Проверка 06_config_merge ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	file := config{Host: "db.internal", Port: 5432}
	c, err := merge(file, nil)
	check(err == nil && c == file, "без флагов → значения из файла (дефолты флагов не затирают)")
	c, err = merge(file, []string{"-port", "6543"})
	check(err == nil && c.Host == "db.internal" && c.Port == 6543, "-port перекрывает только порт")
	c, err = merge(file, []string{"-host", "localhost"})
	check(err == nil && c.Host == "localhost", "явный -host равный дефолту всё равно применяется")
	_, err = merge(file, []string{"-port", "x"})
	check(err != nil, "битый флаг → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
