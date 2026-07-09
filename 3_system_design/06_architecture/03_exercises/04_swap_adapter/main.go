// Упражнение 04: подмена адаптера через порт (хардкор).
//
// Задача: есть порт PaymentGateway. Реализуйте ДВА адаптера:
//  1. AlwaysOK — всегда успешно, возвращает "ok:<amount>".
//  2. AlwaysFail — всегда отказ, возвращает ошибку (Charge -> ("", err != nil)).
//
// И реализуйте usecase (*Checkout).Pay: вызвать gateway.Charge(amount);
// при ошибке вернуть "declined", при успехе — строку от gateway.
//
// Смысл: один и тот же usecase работает поверх разных адаптеров, выбранных
// в composition root, без изменения кода usecase.
//
// Запуск: go run ./03_exercises/04_swap_adapter/main.go
package main

import (
	"errors"
	"fmt"
)

// PaymentGateway — порт (не менять).
type PaymentGateway interface {
	Charge(amount int) (string, error)
}

// Checkout — usecase, зависит от порта (не менять поле).
type Checkout struct {
	gw PaymentGateway
}

func NewCheckout(gw PaymentGateway) *Checkout {
	return &Checkout{gw: gw}
}

// Pay: при ошибке шлюза вернуть "declined", иначе — строку результата шлюза.
func (c *Checkout) Pay(amount int) string {
	// TODO: вызовите c.gw.Charge(amount); обработайте ошибку и успех.
	return ""
}

// AlwaysOK — адаптер, всегда успешный.
type AlwaysOK struct{}

// Charge — заготовка (тест печатает FAIL, пока не реализуете).
// TODO: верните (fmt.Sprintf("ok:%d", amount), nil).
func (AlwaysOK) Charge(amount int) (string, error) {
	return "", nil
}

// AlwaysFail — адаптер, всегда с отказом.
type AlwaysFail struct{}

// Charge — заготовка.
// TODO: верните ("", errors.New("charge failed")) — используйте errCharge.
func (AlwaysFail) Charge(amount int) (string, error) {
	return "", nil
}

var errCharge = errors.New("charge failed")

func main() {
	fmt.Println("--- Проверка 04_swap_adapter ---")

	allOK := true

	// Адаптер 1: успех.
	ok := NewCheckout(&AlwaysOK{})
	if got := ok.Pay(100); got != "ok:100" {
		allOK = false
		fmt.Printf("FAIL: AlwaysOK Pay(100) = %q, ожидалось %q\n", got, "ok:100")
	}

	// Адаптер 2: отказ. Тот же usecase, другой адаптер.
	fail := NewCheckout(&AlwaysFail{})
	if got := fail.Pay(100); got != "declined" {
		allOK = false
		fmt.Printf("FAIL: AlwaysFail Pay(100) = %q, ожидалось %q\n", got, "declined")
	}

	_ = errCharge // подсказка: используйте errors.New в адаптере AlwaysFail

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
