// Упражнение 02: ошибки в таблице (Решение).
package divide

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"exact", 10, 2, 5, nil},
		{"truncated", 7, 2, 3, nil},
		{"negative", -9, 3, -3, nil},
		{"by zero", 1, 0, 0, ErrDivByZero},
	}
	for _, tt := range tests {
		got, err := Divide(tt.a, tt.b)
		if tt.wantErr != nil {
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("%s: err = %v, want %v", tt.name, err, tt.wantErr)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.name, err)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: Divide(%d, %d) = %d, want %d", tt.name, tt.a, tt.b, got, tt.want)
		}
	}
}
