// Упражнение 01: проверка email регулярным выражением.
//
// Задача: реализуйте isValidEmail(s) с помощью пакета regexp:
//   - скомпилируйте шаблон ОДИН раз на уровне пакета (regexp.MustCompile),
//     а не на каждый вызов;
//   - шаблон должен быть заякорен (^...$), иначе "x a@b.io y" пройдёт;
//   - упрощённые правила: локальная часть [a-zA-Z0-9._%+-]+, затем '@',
//     домен [a-zA-Z0-9.-]+, точка и зона из 2+ латинских букв.
//
// Запуск: go run ./03_exercises/01_regex_email/main.go
package main

import (
	"fmt"
	"regexp"
)

// TODO: объявите var emailRe = regexp.MustCompile(`...`)
var _ = regexp.MustCompile // держим импорт живым в заготовке

// isValidEmail сообщает, похожа ли строка на email.
func isValidEmail(s string) bool {
	// TODO: верните emailRe.MatchString(s).
	return true
}

func main() {
	fmt.Println("--- Проверка 01_regex_email ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, ok := range []string{"user@example.com", "first.last+tag@mail.co.uk", "a_b@x.io"} {
		check(isValidEmail(ok), fmt.Sprintf("%q должен быть валидным", ok))
	}
	for _, bad := range []string{"", "plain", "@example.com", "user@", "user@host", "x a@b.io y", "user@domain.c"} {
		check(!isValidEmail(bad), fmt.Sprintf("%q должен быть невалидным", bad))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
