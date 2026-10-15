// Упражнение 02: клиент против httptest.NewServer.
//
// Задача: поднимите настоящий HTTP-сервер на loopback через
// httptest.NewServer (закройте его через t.Cleanup(srv.Close)). Сервер:
//   - /users/42 -> 200 и {"id":42,"name":"Alice"};
//   - /users/500 -> 500;
//   - остальное -> 404.
//
// Тесты: успешный GetUser(42); errors.Is(err, ErrNotFound) для 7;
// ошибка (не ErrNotFound) для 500. Client{BaseURL: srv.URL, HTTP: srv.Client()}.
// Это интеграционный тест клиента: реальные TCP, HTTP и JSON, но без внешних сервисов.
//
// Запуск: go test ./03_exercises/02_httptest_server/
package userclient

import "testing"

func TestGetUser(t *testing.T) {
	t.Fatal("TODO: httptest.NewServer и три сценария")
}
