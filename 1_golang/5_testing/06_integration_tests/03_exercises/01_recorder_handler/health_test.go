// Упражнение 01: handler через httptest.NewRecorder.
//
// Задача: протестируйте Handler без сети:
//   - TestHealthGet: httptest.NewRequest(GET, "/health", nil) +
//     httptest.NewRecorder(); проверьте код 200, заголовок
//     Content-Type == "application/json" и тело (json.Unmarshal -> status == "ok");
//   - TestHealthPost: POST -> 405 и заголовок Allow == "GET".
//
// Запуск: go test ./03_exercises/01_recorder_handler/
package health

import "testing"

func TestHealth(t *testing.T) {
	t.Fatal("TODO: тесты handler через httptest.NewRecorder")
}
