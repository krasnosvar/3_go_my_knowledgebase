// ⏩ Забегание вперёд: пакет unsafe разбирается в 7_under_the_hood/05_reflection_unsafe.
// Эту задачу можно отложить; в обычном коде используйте string(b) и []byte(s).

package main

import (
	"fmt"
	"unsafe"
)

var _ unsafe.Pointer // Держим импорт активным для компиляции заготовки

// BytesToString конвертирует срез байт в строку без копирования и аллокаций.
// Используйте пакет unsafe (функцию unsafe.String, Go 1.20+).
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	// TODO: Реализуйте zero-copy конвертацию []byte -> string.
	return ""
}

// StringToBytes converts string to byte slice without copy and allocations.
// Use unsafe package (Go 1.20+ unsafe.Slice and unsafe.StringData functions).
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	// TODO: Реализуйте zero-copy конвертацию string -> []byte.
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_unsafe_conversions ---")

	b := []byte("hello unsafe")
	s := BytesToString(b)

	b2 := StringToBytes("test")

	t1 := s == "hello unsafe"
	t2 := len(b2) == 4 && string(b2) == "test"

	if t1 && t2 {
		fmt.Println("Тест zero-copy: PASS")
	} else {
		fmt.Printf("Тест zero-copy: FAIL (s=%q, b2=%v)\n", s, b2)
	}
}
