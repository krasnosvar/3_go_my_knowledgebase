// Упражнение 04: spy: проверка вызовов (Решение).
package broadcast

import (
	"errors"
	"slices"
	"testing"
)

type spyNotifier struct {
	calls  []string
	failOn string
}

func (s *spyNotifier) Send(user, _ string) error {
	s.calls = append(s.calls, user)
	if user == s.failOn {
		return errors.New("send failed")
	}
	return nil
}

func TestBroadcastAll(t *testing.T) {
	spy := &spyNotifier{}
	if err := Broadcast(spy, []string{"a", "b", "c"}, "hi"); err != nil {
		t.Fatalf("Broadcast: %v", err)
	}
	if want := []string{"a", "b", "c"}; !slices.Equal(spy.calls, want) {
		t.Errorf("calls = %v, want %v", spy.calls, want)
	}
}

func TestBroadcastStopsOnError(t *testing.T) {
	spy := &spyNotifier{failOn: "b"}
	if err := Broadcast(spy, []string{"a", "b", "c"}, "hi"); err == nil {
		t.Fatal("ожидалась ошибка")
	}
	if want := []string{"a", "b"}; !slices.Equal(spy.calls, want) {
		t.Errorf("calls = %v, want %v", spy.calls, want)
	}
}
