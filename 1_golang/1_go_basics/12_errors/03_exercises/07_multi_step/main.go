package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("invalid request")
	ErrDatabase   = errors.New("db save failure")
)

func step1() error {
	return ErrValidation
}

func step2() error {
	err := step1()
	if err != nil {
		// TODO: Оберните ошибку ErrValidation в контекст "step2: %w"
		return nil
	}
	return nil
}

func step3() error {
	err := step2()
	if err != nil {
		// TODO: Оберните ошибку из step2 в контекст "step3: %w"
		return nil
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_multi_step ---")

	err := step3()

	t1 := err != nil
	t2 := false
	if err != nil {
		t2 = err.Error() == "step3: step2: invalid request"
	}
	t3 := errors.Is(err, ErrValidation)

	if t1 && t2 && t3 {
		fmt.Println("Тест multi-step wrapping: PASS")
	} else {
		fmt.Printf("Тест multi-step wrapping: FAIL (err=%v, isValidation=%t)\n", err, t3)
	}
}
