// Package order — код под тестом для упражнения 02.
package order

import "fmt"

// Payments — внешний платёжный сервис.
type Payments interface {
	Charge(orderID string, amount int) error
}

// Place списывает деньги за заказ; ошибку платежа оборачивает контекстом.
func Place(p Payments, orderID string, amount int) error {
	if err := p.Charge(orderID, amount); err != nil {
		return fmt.Errorf("place order %s: %w", orderID, err)
	}
	return nil
}
