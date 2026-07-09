// Упражнение 01: base62 encode/decode (Решение).
package main

import (
	"fmt"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const base = uint64(len(alphabet)) // 62

func encode(id uint64) string {
	if id == 0 {
		return string(alphabet[0])
	}
	var sb strings.Builder
	for id > 0 {
		sb.WriteByte(alphabet[id%base])
		id /= base
	}
	// Разворачиваем: цифры собирались от младшей к старшей.
	b := []byte(sb.String())
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func decode(code string) uint64 {
	var id uint64
	for i := 0; i < len(code); i++ {
		pos := strings.IndexByte(alphabet, code[i])
		if pos < 0 {
			return 0
		}
		id = id*base + uint64(pos)
	}
	return id
}

func main() {
	fmt.Println("--- Проверка 01_base62 (Решение) ---")

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
