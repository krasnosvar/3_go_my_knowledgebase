package audit

import (
	"slices"
	"testing"
)

type spyLogger struct{ events []string }

func (s *spyLogger) Log(e string) { s.events = append(s.events, e) }

func TestTransferLogsInOrder(t *testing.T) {
	spy := &spyLogger{}
	Transfer(spy, "alice", "bob", 10)
	want := []string{"debit:alice", "credit:bob"}
	if !slices.Equal(spy.events, want) {
		t.Errorf("events = %v, want %v", spy.events, want)
	}
}
