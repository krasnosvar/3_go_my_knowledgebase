// Упражнение 01: fake-репозиторий.
//
// Задача: напишите fakeRepo — структуру, которая реализует Repo и
// запоминает сохранённых пользователей в слайсе saved []User. Затем тесты:
//   - TestRegisterSaves: Register("alice") -> nil, в saved ровно один User{"alice"};
//   - TestRegisterEmptyName: Register("") -> ErrEmptyName, Save НЕ вызывался.
//
// Никаких библиотек моков: fake — обычный тип в тестовом файле.
//
// Запуск: go test ./03_exercises/01_fake_repo/
package users

import "testing"

func TestRegister(t *testing.T) {
	t.Fatal("TODO: fakeRepo и тесты Register")
}
