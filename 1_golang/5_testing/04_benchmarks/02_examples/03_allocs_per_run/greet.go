// Package greet — пример: тест на число аллокаций.
package greet

import "strings"

// Greet собирает приветствие одной аллокацией.
func Greet(name string) string {
	var b strings.Builder
	b.Grow(len("Hello, ") + len(name) + 1)
	b.WriteString("Hello, ")
	b.WriteString(name)
	b.WriteByte('!')
	return b.String()
}
