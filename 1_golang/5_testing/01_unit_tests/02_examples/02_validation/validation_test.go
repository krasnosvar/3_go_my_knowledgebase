package validation

import "testing"

func TestIsAdult(t *testing.T) {
	if !IsAdult(18) {
		t.Fatal("expected true")
	}
}
