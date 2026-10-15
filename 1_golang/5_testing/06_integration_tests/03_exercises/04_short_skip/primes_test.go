// Упражнение 04: медленные тесты и -short.
//
// Задача: напишите два теста:
//   - TestCountPrimesSmall (быстрый, всегда): n = 0, 2, 3, 10 (-> 0, 0, 1, 4);
//   - TestCountPrimesLarge (медленный): n = 50_000_000 -> 3_001_134.
//     В начале — `if testing.Short() { t.Skip("медленный тест: запускайте без -short") }`.
//
// Проверьте оба режима:
//
//	go test -short -v ./03_exercises/04_short_skip/   (Large -> SKIP)
//	go test -v ./03_exercises/04_short_skip/          (Large выполняется)
//
// Так же отделяют интеграционные тесты, которым нужна БД или Docker.
//
// Запуск: go test ./03_exercises/04_short_skip/
package primes

import "testing"

func TestCountPrimes(t *testing.T) {
	t.Fatal("TODO: быстрый тест и медленный с testing.Short()")
}
