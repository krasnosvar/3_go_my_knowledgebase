// Упражнение 08: отмена конвейера через контекст.
//
// Задача: стадии конвейера должны завершаться при отмене ctx, даже если
// никто больше не читает их выход (иначе горутины утекут):
//   - counter(ctx) бесконечно отправляет 1, 2, 3, ...;
//   - square(ctx, in) возводит в квадрат значения из in.
//
// Каждая отправка — через select { case out <- v: ; case <-ctx.Done(): return },
// каждая стадия закрывает свой выход (defer close(out)).
//
// Запуск: go run ./03_exercises/08_pipeline_cancel/main.go
package main

import (
	"context"
	"fmt"
	"time"
)

// counter бесконечно отправляет 1, 2, 3, ... до отмены ctx.
func counter(ctx context.Context) <-chan int {
	// TODO
	return nil
}

// square возводит значения в квадрат до отмены ctx или закрытия in.
func square(ctx context.Context, in <-chan int) <-chan int {
	// TODO
	return nil
}

// --- вспомогательные функции проверки (не меняйте) ---

// collect читает канал до закрытия; ok=false, если не закрылся за d.
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
	fmt.Println("--- Проверка 08_pipeline_cancel ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	out := square(ctx, counter(ctx))

	got := []int{}
	first := time.After(300 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case v := <-out:
			got = append(got, v)
		case <-first:
			i = 3
		}
	}
	check(fmt.Sprint(got) == "[1 4 9]", fmt.Sprintf("первые значения = %v, ожидалось [1 4 9]", got))

	cancel()
	_, closed := collect(out, 300*time.Millisecond)
	check(closed, "после cancel выход конвейера не закрылся — горутины утекли")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
