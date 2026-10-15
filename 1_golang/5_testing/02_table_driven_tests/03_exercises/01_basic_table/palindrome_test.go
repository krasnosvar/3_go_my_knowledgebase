// Упражнение 01: первая табличная проверка.
//
// Задача: напишите TestIsPalindrome в табличном стиле:
//
//	tests := []struct{ name, in string; want bool }{ ... }
//
// и один цикл for _, tt := range tests с t.Errorf.
// Минимум 6 случаев: "", "a", "abba", "Abba", "абба", "go".
//
// Запуск: go test ./03_exercises/01_basic_table/
package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Fatal("TODO: напишите табличный тест для IsPalindrome")
}
