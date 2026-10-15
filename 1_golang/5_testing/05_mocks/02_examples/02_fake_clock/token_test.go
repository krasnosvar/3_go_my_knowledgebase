package token

import (
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	base := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	v := &Validator{now: func() time.Time { return base }} // фиксированное «сейчас»

	if !v.Valid(Token{ExpiresAt: base.Add(time.Minute)}) {
		t.Error("токен, истекающий через минуту, должен быть валиден")
	}
	if v.Valid(Token{ExpiresAt: base}) {
		t.Error("токен, истекающий ровно сейчас, уже невалиден")
	}
}
