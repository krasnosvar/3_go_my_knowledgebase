// Упражнение 05: простой токен в base64.
//
// Задача: реализуйте пару функций:
//   - encodeToken(user, exp) кодирует строку "user:exp" через
//     base64.RawURLEncoding (URL-safe, без '=' на конце);
//   - decodeToken(tok) декодирует токен обратно в (user, exp).
//     Ошибки: невалидный base64, нет ':' (strings.Cut), exp не число.
//
// Это НЕ шифрование и не подпись: base64 читается кем угодно.
//
// Запуск: go run ./03_exercises/05_base64_token/main.go
package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

var _ = base64.RawURLEncoding // держим импорты живыми в заготовке
var _ = strconv.FormatInt
var _ = strings.Cut

// encodeToken кодирует user и время истечения в base64url.
func encodeToken(user string, exp int64) string {
	// TODO: base64.RawURLEncoding.EncodeToString([]byte(user + ":" + exp)).
	return ""
}

// decodeToken разбирает токен, созданный encodeToken.
func decodeToken(tok string) (string, int64, error) {
	// TODO: DecodeString → strings.Cut(s, ":") → strconv.ParseInt.
	return "", 0, nil
}

func main() {
	fmt.Println("--- Проверка 05_base64_token ---")

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
