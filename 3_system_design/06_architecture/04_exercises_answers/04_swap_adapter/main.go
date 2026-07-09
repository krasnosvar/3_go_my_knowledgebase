// Упражнение 04: подмена адаптера через порт (Решение).
package main

import (
	"errors"
	"fmt"
)

type PaymentGateway interface {
	Charge(amount int) (string, error)
}

type Checkout struct {
	gw PaymentGateway
}

func NewCheckout(gw PaymentGateway) *Checkout {
	return &Checkout{gw: gw}
}

func (c *Checkout) Pay(amount int) string {
	res, err := c.gw.Charge(amount)
	if err != nil {
		return "declined"
	}
	return res
}

type AlwaysOK struct{}

func (AlwaysOK) Charge(amount int) (string, error) {
	return fmt.Sprintf("ok:%d", amount), nil
}

type AlwaysFail struct{}

func (AlwaysFail) Charge(amount int) (string, error) {
	return "", errors.New("charge failed")
}

func main() {
	fmt.Println("--- Проверка 04_swap_adapter (Решение) ---")

	allOK := true

	ok := NewCheckout(&AlwaysOK{})
	if got := ok.Pay(100); got != "ok:100" {
		allOK = false
		fmt.Printf("FAIL: AlwaysOK Pay(100) = %q, ожидалось %q\n", got, "ok:100")
	}

	fail := NewCheckout(&AlwaysFail{})
	if got := fail.Pay(100); got != "declined" {
		allOK = false
		fmt.Printf("FAIL: AlwaysFail Pay(100) = %q, ожидалось %q\n", got, "declined")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
