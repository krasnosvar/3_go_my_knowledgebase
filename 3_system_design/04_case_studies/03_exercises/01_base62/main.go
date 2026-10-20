// Упражнение 01: base62 encode/decode.
//
// Задача: реализуйте кодирование инкрементного ID в короткий base62-код
// и обратное декодирование (см. 01_theory/04_example_url_shortener.md).
//
// Алфавит: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// (62 символа). encode(0) == "0". decode должен быть точным обратным encode.
//
// Требование: для всех проверочных id должно выполняться decode(encode(id))
// == id,
// и коды должны совпадать с ожидаемыми.
//
// Запуск: go run ./03_exercises/01_base62/main.go
package main

import (
	"fmt"
	"strings"
)

// Держим импорт strings живым (в заготовке ещё не используется).
// В решении удобен strings.IndexByte / strings.Builder.
var _ = strings.IndexByte

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// encode переводит id в base62-строку. encode(0) == "0".
func encode(id uint64) string {
	// TODO: реализуйте перевод числа в систему счисления с основанием 62.
	// Подсказка: собирайте остатки от деления на 62 и разверните строку.
	return ""
}

// decode переводит base62-строку обратно в число.
func decode(code string) uint64 {
	// TODO: реализуйте свёртку строки в число (id = id*62 + позиция символа).
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_base62 ---")

	// Проверка совпадения кода с ожидаемым.
	enc := []struct {
		id   uint64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{61, "Z"},
		{62, "10"},
		{12345, "3d7"},
	}

	allOK := true
	for _, c := range enc {
		got := encode(c.id)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: encode(%d) = %q, ожидалось %q\n", c.id, got, c.want)
		}
	}

	// Проверка round-trip: decode(encode(id)) == id.
	for _, id := range []uint64{0, 1, 61, 62, 12345, 1_000_000, 3_521_614_606_207} {
		if back := decode(encode(id)); back != id {
			allOK = false
			fmt.Printf("FAIL: round-trip id=%d -> %q -> %d\n", id, encode(id), back)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
