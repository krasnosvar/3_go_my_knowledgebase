// Package divide — код под тестом для упражнения 02.
package divide

import "errors"

// ErrDivByZero — деление на ноль.
var ErrDivByZero = errors.New("divide by zero")

// Divide возвращает a / b (целочисленно) или ErrDivByZero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}
