// Упражнение 02: единые хелперы JSON-ответов.
//
// Задача:
//   - writeJSON(w, status, v): Content-Type "application/json", статус,
//     тело — json-кодирование v;
//   - writeError(w, status, msg): через writeJSON отдать {"error": msg}.
//
// Единый формат ошибок упрощает жизнь клиентам API.
//
// Запуск: go run ./03_exercises/02_json_response
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	// TODO: заголовок, затем WriteHeader, затем json.NewEncoder(w).Encode(v).
}

func writeError(w http.ResponseWriter, status int, msg string) {
	// TODO: используйте writeJSON с map[string]string{"error": msg}.
}

func main() {
	fmt.Println("--- Проверка 02_json_response ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	rec := httptest.NewRecorder()
	writeJSON(rec, 201, map[string]int{"id": 7})
	var got map[string]int
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	check(rec.Code == 201 && got["id"] == 7, "writeJSON: 201 и {\"id\":7}")
	check(rec.Header().Get("Content-Type") == "application/json", "writeJSON ставит Content-Type")

	rec = httptest.NewRecorder()
	writeError(rec, 404, "user not found")
	var e map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &e)
	check(rec.Code == 404 && e["error"] == "user not found", "writeError: 404 и {\"error\":...}")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
