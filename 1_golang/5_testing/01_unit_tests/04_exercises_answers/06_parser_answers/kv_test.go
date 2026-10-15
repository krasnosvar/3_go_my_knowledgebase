// Упражнение 06: тест парсера (Решение).
package kv

import (
	"errors"
	"maps"
	"testing"
)

func TestParseOK(t *testing.T) {
	inputs := map[string]map[string]string{
		"a=1; b = 2": {"a": "1", "b": "2"},
		"":           {},
		"a=1;;b=2":   {"a": "1", "b": "2"},
	}
	for in, want := range inputs {
		got, err := Parse(in)
		if err != nil {
			t.Errorf("Parse(%q) unexpected error: %v", in, err)
			continue
		}
		if !maps.Equal(got, want) {
			t.Errorf("Parse(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestParseBadPair(t *testing.T) {
	for _, in := range []string{"novalue", "=1"} {
		if _, err := Parse(in); !errors.Is(err, ErrBadPair) {
			t.Errorf("Parse(%q) err = %v, want ErrBadPair", in, err)
		}
	}
}
