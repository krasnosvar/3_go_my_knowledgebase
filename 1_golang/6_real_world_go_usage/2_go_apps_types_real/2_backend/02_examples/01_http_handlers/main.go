// Пример: JSON health-handler. В реальном сервисе он регистрируется в mux и
// слушает порт (http.ListenAndServe(":8080", mux)); здесь сервер поднимается
// через httptest на случайном локальном порту, делает запрос и завершается.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type Health struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Health{Status: "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/health")
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, resp.Header.Get("Content-Type"), string(body))
}
