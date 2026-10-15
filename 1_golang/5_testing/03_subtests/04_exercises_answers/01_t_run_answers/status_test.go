// Упражнение 01: подтесты через t.Run (Решение).
package status

import "testing"

func TestClass(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{"continue", 100, "info"},
		{"ok", 200, "success"},
		{"moved", 301, "redirect"},
		{"not found", 404, "client_error"},
		{"unavailable", 503, "server_error"},
		{"below range", 99, "unknown"},
		{"above range", 600, "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Class(tt.code); got != tt.want {
				t.Errorf("Class(%d) = %q, want %q", tt.code, got, tt.want)
			}
		})
	}
}
