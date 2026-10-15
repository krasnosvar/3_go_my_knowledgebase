// Упражнение 01: первая табличная проверка (Решение).
package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", true},
		{"single", "a", true},
		{"even", "abba", true},
		{"mixed case", "Abba", true},
		{"cyrillic", "абба", true},
		{"not palindrome", "go", false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("%s: IsPalindrome(%q) = %t, want %t", tt.name, tt.in, got, tt.want)
		}
	}
}
