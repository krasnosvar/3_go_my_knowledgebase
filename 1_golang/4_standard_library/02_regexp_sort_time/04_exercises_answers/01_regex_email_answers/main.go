// Упражнение 01: проверка email регулярным выражением (Решение).
package main

import (
	"fmt"
	"regexp"
)

var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// isValidEmail сообщает, похожа ли строка на email.
func isValidEmail(s string) bool {
	return emailRe.MatchString(s)
}

func main() {
	fmt.Println("--- Проверка 01_regex_email (Решение) ---")

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
