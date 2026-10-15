// Упражнение 07: граница с внешним миром (Решение).
package words

import (
	"errors"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	got, err := CountWords(strings.NewReader("go is  fun\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 3 {
		t.Errorf("CountWords = %d, want 3", got)
	}
}

var errBoom = errors.New("boom")

// errReader — Reader, который всегда возвращает ошибку.
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errBoom }

func TestCountWordsReadError(t *testing.T) {
	if _, err := CountWords(errReader{}); !errors.Is(err, errBoom) {
		t.Errorf("err = %v, want errBoom", err)
	}
}
