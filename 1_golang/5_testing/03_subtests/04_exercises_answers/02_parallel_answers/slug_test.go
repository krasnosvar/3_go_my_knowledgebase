// Упражнение 02: параллельные подтесты (Решение).
package slug

import "testing"

func TestSlugify(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, in, want string
	}{
		{"latin", "Hello World", "hello-world"},
		{"extra spaces and dots", "  Go   1.22 ", "go-122"},
		{"cyrillic", "Привет, мир!", "привет-мир"},
		{"only punctuation", "!!!", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // подтест ставится на паузу и продолжается вместе с остальными
			if got := Slugify(tt.in); got != tt.want {
				t.Errorf("Slugify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
