// Упражнение 04: проверка кворума W + R > N (Решение).
package main

import (
	"errors"
	"fmt"
)

var _ = errors.New

func quorumConsistent(n, w, r int) (bool, error) {
	if w < 1 || w > n || r < 1 || r > n {
		return false, fmt.Errorf("невалидные параметры: n=%d w=%d r=%d (нужно 1<=w<=n и 1<=r<=n)", n, w, r)
	}
	return w+r > n, nil
}

func main() {
	fmt.Println("--- Проверка 04_check_quorum (Решение) ---")

	type tc struct {
		n, w, r   int
		want      bool
		wantError bool
	}
	cases := []tc{
		{3, 2, 2, true, false},
		{3, 1, 1, false, false},
		{3, 3, 1, true, false},
		{5, 3, 3, true, false},
		{5, 2, 2, false, false},
		{3, 0, 2, false, true},
		{3, 2, 4, false, true},
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
