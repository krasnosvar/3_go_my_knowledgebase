package split

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"simple", "a b", []string{"a", "b"}},
	}
	for _, tt := range tests {
		if got := Words(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%s: got %v want %v", tt.name, got, tt.want)
		}
	}
}
