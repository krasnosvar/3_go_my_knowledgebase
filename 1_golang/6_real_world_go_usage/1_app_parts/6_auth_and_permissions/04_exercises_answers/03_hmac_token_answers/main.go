// Упражнение 03: подписанный токен на HMAC-SHA256 (Решение).
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

func verify(secret []byte, token string, now int64) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errMalformed
	}
	want := mac(secret, parts[0]+"."+parts[1])
	if !hmac.Equal([]byte(want), []byte(parts[2])) {
		return "", errBadSig
	}
	exp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return "", errMalformed
	}
	if now >= exp {
		return "", errExpired
	}
	return parts[0], nil
}

func main() {
	fmt.Println("--- Проверка 03_hmac_token (Решение) ---")

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
