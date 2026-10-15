// Package palindrome — код под тестом для упражнения 01.
package palindrome

import "strings"

// IsPalindrome сообщает, читается ли s одинаково в обе стороны.
// Регистр не учитывается, сравнение идёт по рунам (работает с кириллицей).
func IsPalindrome(s string) bool {
	r := []rune(strings.ToLower(s))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
