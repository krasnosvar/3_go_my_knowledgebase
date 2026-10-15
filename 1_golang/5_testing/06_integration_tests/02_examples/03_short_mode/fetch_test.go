package fetch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -short ./...  — тест пропускается (быстрый прогон в CI на каждый коммит)
// go test ./...         — выполняется полностью
func TestBodyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("интеграционный тест: запускайте без -short")
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}))
	t.Cleanup(srv.Close)

	got, err := Body(srv.Client(), srv.URL+"/ping")
	if err != nil {
		t.Fatalf("Body: %v", err)
	}
	if got != "pong" {
		t.Errorf("Body = %q, want pong", got)
	}
}
