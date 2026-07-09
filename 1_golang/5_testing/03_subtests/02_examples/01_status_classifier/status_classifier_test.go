package status

import "testing"

func TestClassify(t *testing.T) {
	t.Run("server", func(t *testing.T) {
		if got := Classify(500); got != "server" {
			t.Fatalf("got %s", got)
		}
	})
	t.Run("client", func(t *testing.T) {
		if got := Classify(404); got != "client" {
			t.Fatalf("got %s", got)
		}
	})
}
