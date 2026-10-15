// Package port — пример: табличный тест с ожидаемой ошибкой.
package port

import (
	"errors"
	"strconv"
)

// ErrInvalidPort — строка не является портом 1..65535.
var ErrInvalidPort = errors.New("invalid port")

// ParsePort разбирает номер TCP-порта.
func ParsePort(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 || n > 65535 {
		return 0, ErrInvalidPort
	}
	return n, nil
}
