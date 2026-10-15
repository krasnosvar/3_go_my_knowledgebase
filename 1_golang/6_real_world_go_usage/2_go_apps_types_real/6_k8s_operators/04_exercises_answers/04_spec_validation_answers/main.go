// Упражнение 04: валидация spec кастомного ресурса (Решение).
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

func validate(s spec) error {
	var errs []error
	if s.Replicas < 0 || s.Replicas > 10 {
		errs = append(errs, errors.New("spec.replicas: must be 0..10"))
	}
	name := s.Image[strings.LastIndex(s.Image, "/")+1:]
	_, tag, ok := strings.Cut(name, ":")
	switch {
	case s.Image == "" || !ok || tag == "":
		errs = append(errs, errors.New("spec.image: tag is required"))
	case tag == "latest":
		errs = append(errs, errors.New("spec.image: latest tag is not allowed"))
	}
	return errors.Join(errs...)
}

func main() {
	fmt.Println("--- Проверка 04_spec_validation (Решение) ---")

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
