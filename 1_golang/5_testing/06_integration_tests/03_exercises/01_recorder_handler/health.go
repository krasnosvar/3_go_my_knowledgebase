// Package health — код под тестом для упражнения 01.
package health

import (
	"encoding/json"
	"net/http"
)

// Handler отвечает на GET JSON-ом {"status":"ok"}; другие методы — 405.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
