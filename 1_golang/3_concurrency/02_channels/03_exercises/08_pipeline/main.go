// Упражнение 08: pipeline из двух стадий.
//
// Задача: соберите конвейер из стадий, каждая — горутина со своим
// выходным каналом, который она закрывает по завершении входа:
//   - multiply(in, k) — умножает каждое значение на k;
//   - filterGreater(in, min) — пропускает только значения > min.
//
// Источник gen уже написан. Цепочка gen(1..5) → ×3 → >6 даёт [9 12 15].
//
// Запуск: go run ./03_exercises/08_pipeline/main.go
package main

import (
	"fmt"
	"time"
)

// gen отправляет nums в канал и закрывает его (готово).
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// multiply умножает каждое значение из in на k.
func multiply(in <-chan int, k int) <-chan int {
	// TODO: out-канал, горутина с range по in, close(out) в конце.
	return nil
}

// filterGreater пропускает значения строго больше min.
func filterGreater(in <-chan int, min int) <-chan int {
	// TODO: аналогично multiply, но с фильтром.
	return nil
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если канал не закрылся за d.
func collect(ch <-chan int, d time.Duration) ([]int, bool) {
	res := []int{}
	deadline := time.After(d)
	for {
		select {
		case v, open := <-ch:
			if !open {
				return res, true
			}
			res = append(res, v)
		case <-deadline:
			return res, false
		}
	}
}

func main() {
	fmt.Println("--- Проверка 08_pipeline ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got, closed := collect(filterGreater(multiply(gen(1, 2, 3, 4, 5), 3), 6), 300*time.Millisecond)
	check(closed, "выходной канал конвейера не закрылся")
	check(fmt.Sprint(got) == "[9 12 15]", fmt.Sprintf("конвейер отдал %v, ожидалось [9 12 15]", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
