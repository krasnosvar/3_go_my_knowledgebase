// Упражнение 05: простой токен в base64 (Решение).
package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// encodeToken кодирует user и время истечения в base64url.
func encodeToken(user string, exp int64) string {
	raw := user + ":" + strconv.FormatInt(exp, 10)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

// decodeToken разбирает токен, созданный encodeToken.
func decodeToken(tok string) (string, int64, error) {
	raw, err := base64.RawURLEncoding.DecodeString(tok)
	if err != nil {
		return "", 0, fmt.Errorf("decode base64: %w", err)
	}
	user, expStr, ok := strings.Cut(string(raw), ":")
	if !ok {
		return "", 0, fmt.Errorf("token %q: missing ':'", raw)
	}
	exp, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("token exp %q: %w", expStr, err)
	}
	return user, exp, nil
}

func main() {
	fmt.Println("--- Проверка 05_base64_token (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	tok := encodeToken("ann", 1700000000)
	check(tok == "YW5uOjE3MDAwMDAwMDA", fmt.Sprintf("encodeToken = %q, ожидалось \"YW5uOjE3MDAwMDAwMDA\"", tok))
	check(!strings.ContainsAny(tok, "=+/"), "RawURLEncoding не содержит '=', '+', '/'")

	user, exp, err := decodeToken(tok)
	check(err == nil && user == "ann" && exp == 1700000000, fmt.Sprintf("decodeToken = %q, %d, %v", user, exp, err))

	_, _, err = decodeToken("!!!")
	check(err != nil, "невалидный base64 должен давать ошибку")
	_, _, err = decodeToken(base64.RawURLEncoding.EncodeToString([]byte("no-colon")))
	check(err != nil, "токен без ':' должен давать ошибку")
	_, _, err = decodeToken(base64.RawURLEncoding.EncodeToString([]byte("ann:soon")))
	check(err != nil, "нечисловой exp должен давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
