// Package kv — код под тестом для упражнения 06.
package kv

import (
	"errors"
	"fmt"
	"strings"
)

// ErrBadPair — пара не в формате key=value или с пустым ключом.
var ErrBadPair = errors.New("kv: bad pair")

// Parse разбирает строку вида "a=1; b=2" в map. Пробелы вокруг ключей и
// значений обрезаются, пустые сегменты (";;") пропускаются.
func Parse(s string) (map[string]string, error) {
	out := make(map[string]string)
	for _, part := range strings.Split(s, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		k, v, ok := strings.Cut(part, "=")
		k = strings.TrimSpace(k)
		if !ok || k == "" {
			return nil, fmt.Errorf("%w: %q", ErrBadPair, part)
		}
		out[k] = strings.TrimSpace(v)
	}
	return out, nil
}
