// Упражнение 04: проверка кворума W + R > N (Хардкор).
//
// Задача: реализуйте quorumConsistent(n, w, r) — гарантирует ли данная
// конфигурация кворума согласованное чтение.
//
// Правила:
//   - параметры валидны, если 1 <= w <= n и 1 <= r <= n (иначе -> ошибка);
//   - согласованность гарантирована тогда и только тогда, когда w + r > n
//     (множества записи и чтения обязаны пересекаться).
//
// Сигнатура: quorumConsistent(n, w, r int) (bool, error).
// При невалидных параметрах верните (false, ошибка != nil).
//
// Запуск: go run ./03_exercises/04_check_quorum/main.go
package main

import (
	"errors"
	"fmt"
)

// Держим импорт errors живым в заготовке (в решении он используется).
var _ = errors.New

// quorumConsistent сообщает, даёт ли конфигурация (n,w,r) согласованное чтение.
func quorumConsistent(n, w, r int) (bool, error) {
	// TODO:
	//   1) проверьте валидность: 1 <= w <= n и 1 <= r <= n, иначе верните ошибку;
	//   2) верните (w+r > n, nil).
	return false, nil
}

func main() {
	fmt.Println("--- Проверка 04_check_quorum ---")

	type tc struct {
		n, w, r   int
		want      bool
		wantError bool
	}
	cases := []tc{
		{3, 2, 2, true, false},  // 4 > 3
		{3, 1, 1, false, false}, // 2 <= 3
		{3, 3, 1, true, false},  // 4 > 3
		{5, 3, 3, true, false},  // 6 > 5
		{5, 2, 2, false, false}, // 4 <= 5
		{3, 0, 2, false, true},  // w < 1 -> ошибка
		{3, 2, 4, false, true},  // r > n -> ошибка
	}

	allOK := true
	for _, c := range cases {
		got, err := quorumConsistent(c.n, c.w, c.r)
		if c.wantError {
			if err == nil {
				allOK = false
				fmt.Printf("FAIL: quorumConsistent(%d,%d,%d) — ожидалась ошибка, её нет\n", c.n, c.w, c.r)
			}
			continue
		}
		if err != nil {
			allOK = false
			fmt.Printf("FAIL: quorumConsistent(%d,%d,%d) вернул неожиданную ошибку: %v\n", c.n, c.w, c.r, err)
			continue
		}
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: quorumConsistent(%d,%d,%d) = %v, ожидалось %v\n", c.n, c.w, c.r, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
