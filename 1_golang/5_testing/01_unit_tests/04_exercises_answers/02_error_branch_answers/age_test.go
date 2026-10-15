// Упражнение 02: тест веток ошибок (Решение).
package age

import (
	"errors"
	"testing"
)

func TestParseAgeOK(t *testing.T) {
	got, err := ParseAge("42")
	if err != nil {
		t.Fatalf("ParseAge(\"42\") unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("ParseAge(\"42\") = %d, want 42", got)
	}
}

func TestParseAgeNotNumber(t *testing.T) {
	_, err := ParseAge("abc")
	if !errors.Is(err, ErrNotNumber) {
		t.Errorf("ParseAge(\"abc\") err = %v, want ErrNotNumber", err)
	}
}

func TestParseAgeOutOfRange(t *testing.T) {
	for _, in := range []string{"-1", "151"} {
		if _, err := ParseAge(in); !errors.Is(err, ErrOutOfRange) {
			t.Errorf("ParseAge(%q) err = %v, want ErrOutOfRange", in, err)
		}
	}
}
