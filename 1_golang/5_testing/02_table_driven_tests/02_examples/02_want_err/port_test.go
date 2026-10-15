package port

import (
	"errors"
	"testing"
)

func TestParsePort(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int
		wantErr error
	}{
		{"http", "80", 80, nil},
		{"max", "65535", 65535, nil},
		{"zero", "0", 0, ErrInvalidPort},
		{"too big", "65536", 0, ErrInvalidPort},
		{"not a number", "http", 0, ErrInvalidPort},
	}
	for _, tt := range tests {
		got, err := ParsePort(tt.in)
		if !errors.Is(err, tt.wantErr) { // errors.Is(nil, nil) == true
			t.Errorf("%s: err = %v, want %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%s: ParsePort(%q) = %d, want %d", tt.name, tt.in, got, tt.want)
		}
	}
}
