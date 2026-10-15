// Упражнение 02: внедрение ошибки через fake (Решение).
package order

import (
	"errors"
	"strings"
	"testing"
)

type fakePayments struct {
	err error // что вернуть из Charge
}

func (f fakePayments) Charge(string, int) error { return f.err }

var errDeclined = errors.New("card declined")

func TestPlaceOK(t *testing.T) {
	if err := Place(fakePayments{}, "A-1", 100); err != nil {
		t.Errorf("Place: unexpected error %v", err)
	}
}

func TestPlaceChargeFails(t *testing.T) {
	err := Place(fakePayments{err: errDeclined}, "A-2", 100)
	if !errors.Is(err, errDeclined) {
		t.Fatalf("err = %v, want wrapped errDeclined", err)
	}
	if !strings.Contains(err.Error(), "A-2") {
		t.Errorf("в ошибке нет id заказа: %v", err)
	}
}
