package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	r := httptest.NewRecorder()
	health(r, httptest.NewRequest(http.MethodGet, "/health", nil))
	fmt.Println(r.Code, r.Body.String() != "")
}
