// Упражнение 03: нулевое значение типа (Решение).
package stack

import "testing"

func TestStackZeroValue(t *testing.T) {
	var s Stack
	if s.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", s.Len())
	}
	if _, ok := s.Pop(); ok {
		t.Fatal("Pop() на пустом стеке должен вернуть ok == false")
	}

	s.Push(1)
	s.Push(2)
	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("первый Pop() = (%d, %t), want (2, true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("второй Pop() = (%d, %t), want (1, true)", v, ok)
	}
	if s.Len() != 0 {
		t.Errorf("после двух Pop Len() = %d, want 0", s.Len())
	}
}
