package parse

import "testing"

func TestParsePositive(t *testing.T) {
	if _, err := ParsePositive("-1"); err == nil {
		t.Fatal("expected error")
	}
}
