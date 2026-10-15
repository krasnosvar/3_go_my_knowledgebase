// Упражнение 02: обработка неудачного assertion.
//
// Задача: реализуйте mustString так, чтобы она НИКОГДА не паниковала.
// Однозначная форма v.(string) паникует при несовпадении типа — используйте
// форму с ok. При несовпадении верните ошибку с текстом
// "ожидалась string, получено <тип>" (тип — через %T, для nil это "<nil>").
//
// Запуск: go run ./03_exercises/02_assertion_fail/main.go
package main

import (
	"fmt"
)

// mustString возвращает строку или ошибку, без паники.
func mustString(v any) (string, error) {
	// TODO: s, ok := v.(string); при !ok — fmt.Errorf("ожидалась string, получено %T", v).
	return "", nil
}

func main() {
	fmt.Println("--- Проверка 02_assertion_fail ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	safe := func(v any) (s string, err error, panicked bool) {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()
		s, err = mustString(v)
		return
	}

	s, err, p := safe("ok")
	check(!p && err == nil && s == "ok", fmt.Sprintf("mustString(\"ok\"): получено (%q, %v, panic=%t)", s, err, p))

	cases := []struct {
		in   any
		want string
	}{
		{42, "ожидалась string, получено int"},
		{3.5, "ожидалась string, получено float64"},
		{nil, "ожидалась string, получено <nil>"},
	}
	for _, c := range cases {
		_, err, p := safe(c.in)
		check(!p, fmt.Sprintf("mustString(%#v) не должна паниковать", c.in))
		check(err != nil && err.Error() == c.want, fmt.Sprintf("mustString(%#v): ожидалась ошибка %q, получено %v", c.in, c.want, err))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
