// Пример 01: base62-кодирование инкрементного ID для сокращателя ссылок.
//
// Идея (см. 01_theory/04_example_url_shortener.md): вместо хэширования URL
// берём глобально уникальный инкрементный ID и кодируем его в base62
// (алфавит [0-9a-zA-Z], безопасный для URL). Биекция число<->строка
// гарантирует уникальность коротких кодов без коллизий.
//
// Программа кодирует несколько ID в короткие коды, декодирует обратно,
// проверяет round-trip и завершается.
//
// Запуск: go run ./02_examples/01_base62_shortener/main.go
package main

import (
	"fmt"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const base = uint64(len(alphabet)) // 62

// encode переводит число id в base62-строку (короткий код).
// id == 0 кодируется как "0".
func encode(id uint64) string {
	if id == 0 {
		return string(alphabet[0])
	}
	// Собираем цифры от младшей к старшей, затем разворачиваем.
	var sb strings.Builder
	for id > 0 {
		rem := id % base
		sb.WriteByte(alphabet[rem])
		id /= base
	}
	return reverse(sb.String())
}

// decode переводит base62-строку обратно в число.
// Возвращает ошибку, если встретился символ вне алфавита.
func decode(code string) (uint64, error) {
	var id uint64
	for i := 0; i < len(code); i++ {
		pos := strings.IndexByte(alphabet, code[i])
		if pos < 0 {
			return 0, fmt.Errorf("недопустимый символ %q в коде %q", code[i], code)
		}
		id = id*base + uint64(pos)
	}
	return id, nil
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	fmt.Println("--- base62-сокращатель: id <-> код ---")

	// Диапазон base62 на длину кода (ориентир из теории):
	// 62^7 ≈ 3.5e12 — хватает на сотни миллиардов ссылок.
	fmt.Printf("Ёмкость 7 символов: 62^7 ≈ %.2e значений\n\n", pow62(7))

	ids := []uint64{0, 1, 61, 62, 12345, 1_000_000, 3_521_614_606_207} // последний ≈ 62^7-1

	for _, id := range ids {
		code := encode(id)
		back, err := decode(code)
		if err != nil {
			fmt.Printf("id=%d -> код=%q : ОШИБКА декодирования: %v\n", id, code, err)
			continue
		}
		status := "ok"
		if back != id {
			status = "MISMATCH"
		}
		fmt.Printf("id=%-15d -> код=%-8q -> id=%-15d [%s]\n", id, code, back, status)
	}

	fmt.Println("\nГотово, программа завершается.")
}

// pow62 возвращает 62^n как float64 (для печати ёмкости).
func pow62(n int) float64 {
	res := 1.0
	for i := 0; i < n; i++ {
		res *= 62
	}
	return res
}
