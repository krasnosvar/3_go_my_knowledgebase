// Упражнение 06: слияние конфига из файла и флагов (Решение).
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

func merge(file config, args []string) (config, error) {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	host := fs.String("host", "localhost", "host")
	port := fs.Int("port", 80, "port")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	cfg := file
	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "host":
			cfg.Host = *host
		case "port":
			cfg.Port = *port
		}
	})
	return cfg, nil
}

func main() {
	fmt.Println("--- Проверка 06_config_merge (Решение) ---")

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
