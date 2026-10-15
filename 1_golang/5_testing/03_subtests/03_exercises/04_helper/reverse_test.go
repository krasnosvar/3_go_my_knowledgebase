// Упражнение 04: тестовый хелпер с t.Helper.
//
// Задача:
//  1. Напишите хелпер assertEqual(t *testing.T, got, want string), который
//     вызывает t.Helper() и при несовпадении — t.Errorf.
//  2. Напишите TestReverse с подтестами "ascii" ("abc" -> "cba"),
//     "cyrillic" ("привет" -> "тевирп"), "empty" ("" -> ""), "twice"
//     (Reverse(Reverse(s)) == s) — все проверки через assertEqual.
//
// Благодаря t.Helper() сообщение об ошибке укажет строку в подтесте,
// а не внутри хелпера — проверьте, временно сломав ожидание.
//
// Запуск: go test ./03_exercises/04_helper/
package reverse

import "testing"

func TestReverse(t *testing.T) {
	t.Fatal("TODO: хелпер assertEqual и подтесты")
}
