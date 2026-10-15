// Package concat — код для упражнения 02 (Решение).
package concat

import "strings"

// Join склеивает parts через sep одной аллокацией.
func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	n := len(sep) * (len(parts) - 1)
	for _, p := range parts {
		n += len(p)
	}
	var b strings.Builder
	b.Grow(n) // единственная аллокация: буфер сразу нужного размера
	for i, p := range parts {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(p)
	}
	return b.String() // Builder отдаёт строку без копирования
}
