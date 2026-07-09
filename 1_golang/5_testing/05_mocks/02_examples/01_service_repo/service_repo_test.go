package svc

import "testing"

type fakeRepo struct{ called bool }

func (f *fakeRepo) Save(name string) error { f.called = true; return nil }

func TestCreateCallsRepo(t *testing.T) {
	f := &fakeRepo{}
	s := New(f)
	if err := s.Create("den"); err != nil {
		t.Fatal(err)
	}
	if !f.called {
		t.Fatal("expected Save to be called")
	}
}
