package main

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Health{Status: "ok"})
}

func main() {
	http.HandleFunc("/health", health)
	_ = http.ListenAndServe(":8080", nil)
}
