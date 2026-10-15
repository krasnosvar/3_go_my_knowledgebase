// Package age — код под тестом для упражнения 02.
package age

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNotNumber — строка не является целым числом.
	ErrNotNumber = errors.New("age: not a number")
	// ErrOutOfRange — число вне диапазона 0..150.
	ErrOutOfRange = errors.New("age: out of range")
)

// ParseAge разбирает возраст из строки. Допустимы значения 0..150.
func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrNotNumber, s)
	}
	if n < 0 || n > 150 {
		return 0, fmt.Errorf("%w: %d", ErrOutOfRange, n)
	}
	return n, nil
}
