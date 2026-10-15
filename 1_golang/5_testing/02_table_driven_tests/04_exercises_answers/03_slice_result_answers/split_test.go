// Упражнение 03: сравнение слайсов в таблице (Решение).
package split

import (
	"slices"
	"testing"
)

func TestSplitTrim(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"simple", "a,b,c", []string{"a", "b", "c"}},
		{"spaces", " a , b ", []string{"a", "b"}},
		{"empty part", "a,,b", []string{"a", "b"}},
		{"empty input", "", nil},
		{"only separators", ",,,", nil},
	}
	for _, tt := range tests {
		if got := SplitTrim(tt.in, ","); !slices.Equal(got, tt.want) {
			t.Errorf("%s: SplitTrim(%q) = %q, want %q", tt.name, tt.in, got, tt.want)
		}
	}
}
