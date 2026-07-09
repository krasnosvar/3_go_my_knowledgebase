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
		return fmt.Errorf("step2: %w", err)
	}
	return nil
}

func step3() error {
	err := step2()
	if err != nil {
		return fmt.Errorf("step3: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_multi_step (Решение) ---")

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
