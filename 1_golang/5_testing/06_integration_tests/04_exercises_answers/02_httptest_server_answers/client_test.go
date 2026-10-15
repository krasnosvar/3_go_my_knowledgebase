// Упражнение 02: клиент против httptest.NewServer (Решение).
package userclient

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(t *testing.T) *Client {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/users/42", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":42,"name":"Alice"}`))
	})
	mux.HandleFunc("/users/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux) // неизвестные пути mux отдаёт как 404
	t.Cleanup(srv.Close)
	return &Client{BaseURL: srv.URL, HTTP: srv.Client()}
}

func TestGetUser(t *testing.T) {
	c := newTestClient(t)
	ctx := context.Background()

	t.Run("found", func(t *testing.T) {
		u, err := c.GetUser(ctx, 42)
		if err != nil {
			t.Fatalf("GetUser: %v", err)
		}
		if u.ID != 42 || u.Name != "Alice" {
			t.Errorf("user = %+v", u)
		}
	})
	t.Run("not found", func(t *testing.T) {
		if _, err := c.GetUser(ctx, 7); !errors.Is(err, ErrNotFound) {
			t.Errorf("err = %v, want ErrNotFound", err)
		}
	})
	t.Run("server error", func(t *testing.T) {
		_, err := c.GetUser(ctx, 500)
		if err == nil || errors.Is(err, ErrNotFound) {
			t.Errorf("err = %v, want unexpected status error", err)
		}
	})
}
