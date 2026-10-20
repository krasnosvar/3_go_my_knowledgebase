// Упражнение 12: обернуть вызов таймаутом через context.
//
// Задача: реализуйте callWithTimeout — запустить work в отдельной
// горутине и вернуть её результат, но не ждать дольше, чем позволяет
// ctx (context с дедлайном).
//
//   - если work успела вернуть результат до отмены ctx -> вернуть
//     (result, nil);
//   - если ctx отменился (дедлайн истёк) раньше -> вернуть (0, ctx.Err()).
//
// Подсказка: запустите work в горутине, шлющей результат в канал;
// затем select по этому каналу и по ctx.Done().
//
// Детерминизм: в тесте «быстрая» work завершается сразу, «медленная»
// не завершится никогда (канал, который никто не закрывает) — так что
// сработает именно таймаут context, без гонок по реальному времени.
//
// Запуск: go run ./03_exercises/12_call_with_timeout/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// callWithTimeout выполняет work с учётом дедлайна ctx.
func callWithTimeout(ctx context.Context, work func() int) (int, error) {
	// TODO: запустить work в горутине, вернуть результат либо ctx.Err()
	// по select между результатом и ctx.Done().
	return 0, errors.New("не реализовано")
}

func main() {
	fmt.Println("--- Проверка 12_call_with_timeout ---")

	allOK := true

	// Случай 1: быстрая работа успевает — ждём (42, nil).
	// ctx уже отменён СРАЗУ, но work завершилась мгновенно; чтобы тест был
	// детерминированным, используем не-отменённый ctx с запасом.
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Hour)
	defer cancel1()
	got, err := callWithTimeout(ctx1, func() int { return 42 })
	if err != nil || got != 42 {
		allOK = false
		fmt.Printf("FAIL: быстрая work -> (%d, %v), ожидалось (42, nil)\n", got, err)
	}

	// Случай 2: work никогда не завершается -> должен сработать таймаут.
	// Отменяем ctx заранее, чтобы select ушёл в ctx.Done() без ожидания.
	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2() // ctx уже отменён
	never := func() int {
		<-make(chan struct{}) // блокируется навсегда
		return 1
	}
	got, err = callWithTimeout(ctx2, never)
	if err == nil {
		allOK = false
		fmt.Printf("FAIL: медленная work -> (%d, nil), ожидалась ошибка таймаута\n", got)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
