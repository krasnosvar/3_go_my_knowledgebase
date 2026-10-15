// Упражнение 03: подписанный токен (идея JWT на HMAC-SHA256).
//
// Формат токена: "<userID>.<exp>.<sig>", где exp — unix-время истечения,
// sig = hex(HMAC-SHA256(secret, "<userID>.<exp>")).
//
// Задача: реализуйте verify(secret []byte, token string, now int64) (string, error):
//   - неверное число частей → errMalformed;
//   - подпись не совпала → errBadSig (сравнение через hmac.Equal — constant-time);
//   - now >= exp → errExpired;
//   - иначе вернуть userID.
//
// Функция sign уже написана — посмотрите, как считается подпись.
//
// Запуск: go run ./03_exercises/03_hmac_token
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errMalformed = errors.New("malformed token")
	errBadSig    = errors.New("bad signature")
	errExpired   = errors.New("token expired")
)

func mac(secret []byte, payload string) string {
	m := hmac.New(sha256.New, secret)
	m.Write([]byte(payload))
	return hex.EncodeToString(m.Sum(nil))
}

func sign(secret []byte, userID string, exp int64) string {
	payload := userID + "." + strconv.FormatInt(exp, 10)
	return payload + "." + mac(secret, payload)
}

var _ = strings.Split

func verify(secret []byte, token string, now int64) (string, error) {
	// TODO: разбейте токен, проверьте подпись (hmac.Equal), затем срок.
	return "", nil
}

func main() {
	fmt.Println("--- Проверка 03_hmac_token ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	secret := []byte("s3cret")
	tok := sign(secret, "u1", 1000)

	uid, err := verify(secret, tok, 999)
	check(err == nil && uid == "u1", "валидный токен → u1")
	_, err = verify(secret, tok, 1000)
	check(errors.Is(err, errExpired), "now == exp → errExpired")
	_, err = verify([]byte("other"), tok, 1)
	check(errors.Is(err, errBadSig), "чужой секрет → errBadSig")
	forged := strings.Replace(tok, "u1.", "admin.", 1)
	_, err = verify(secret, forged, 1)
	check(errors.Is(err, errBadSig), "подменённый userID → errBadSig")
	_, err = verify(secret, "garbage", 1)
	check(errors.Is(err, errMalformed), "мусор → errMalformed")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
