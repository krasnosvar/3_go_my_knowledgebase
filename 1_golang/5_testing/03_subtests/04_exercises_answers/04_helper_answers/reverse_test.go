// Упражнение 04: тестовый хелпер с t.Helper (Решение).
package reverse

import "testing"

func assertEqual(t *testing.T, got, want string) {
	t.Helper() // строка ошибки будет указывать на место вызова
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverse(t *testing.T) {
	t.Run("ascii", func(t *testing.T) { assertEqual(t, Reverse("abc"), "cba") })
	t.Run("cyrillic", func(t *testing.T) { assertEqual(t, Reverse("привет"), "тевирп") })
	t.Run("empty", func(t *testing.T) { assertEqual(t, Reverse(""), "") })
	t.Run("twice", func(t *testing.T) {
		s := "Go, 世界!"
		assertEqual(t, Reverse(Reverse(s)), s)
	})
}
