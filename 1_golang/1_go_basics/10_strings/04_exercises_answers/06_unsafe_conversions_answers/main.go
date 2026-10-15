// ⏩ Забегание вперёд: пакет unsafe разбирается в 7_under_the_hood/05_reflection_unsafe.

package main

import (
	"fmt"
	"unsafe"
)

func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func main() {
	fmt.Println("--- Проверка 06_unsafe_conversions (Решение) ---")

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
