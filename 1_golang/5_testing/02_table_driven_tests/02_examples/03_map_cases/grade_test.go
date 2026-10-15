package grade

import "testing"

func TestGrade(t *testing.T) {
	tests := map[string]struct {
		score int
		want  string
	}{
		"top":          {100, "A"},
		"A boundary":   {90, "A"},
		"just below A": {89, "B"},
		"C boundary":   {60, "C"},
		"fail":         {59, "F"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Grade(tt.score); got != tt.want {
				t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}
