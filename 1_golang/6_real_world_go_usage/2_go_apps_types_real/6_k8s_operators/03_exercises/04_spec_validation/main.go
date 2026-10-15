// Упражнение 04: валидация spec кастомного ресурса.
//
// Задача: реализуйте validate(s spec) error (как validating webhook):
//   - Replicas в диапазоне 0..10 → иначе "spec.replicas: must be 0..10";
//   - Image не пустой и содержит тег (":" после последнего "/") →
//     иначе "spec.image: tag is required";
//   - тег "latest" запрещён → "spec.image: latest tag is not allowed";
//   - вернуть все нарушения через errors.Join; nil, если всё ок.
//
// Запуск: go run ./03_exercises/04_spec_validation
package main

import (
	"errors"
	"fmt"
	"strings"
)

type spec struct {
	Replicas int
	Image    string
}

var _ = errors.Join
var _ = strings.LastIndex

func validate(s spec) error {
	// TODO: проверки replicas и image, errors.Join.
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_spec_validation ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	msg := func(err error) string {
		if err == nil {
			return ""
		}
		return err.Error()
	}

	check(validate(spec{Replicas: 3, Image: "registry:5000/team/app:1.2.3"}) == nil, "валидный spec (порт registry не путается с тегом)")
	check(strings.Contains(msg(validate(spec{Replicas: 3, Image: "registry:5000/team/app"})), "tag is required"), "без тега → ошибка")
	check(strings.Contains(msg(validate(spec{Replicas: 1, Image: "nginx:latest"})), "latest tag"), "latest запрещён")
	m := msg(validate(spec{Replicas: 42}))
	check(strings.Contains(m, "spec.replicas") && strings.Contains(m, "spec.image"), "обе ошибки сразу: "+m)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
