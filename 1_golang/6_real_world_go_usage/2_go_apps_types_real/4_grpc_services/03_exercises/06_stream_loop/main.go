// Упражнение 06: server-streaming цикл.
//
// В gRPC сервер стримит ответы через stream.Send(msg); stream.Context()
// отменяется, когда клиент ушёл.
//
// Задача: реализуйте streamAll(ctx context.Context, items []string, s sender) (int, error):
//   - отправлять элементы по одному через s.Send;
//   - перед каждой отправкой проверять ctx.Err() — при отмене вернуть
//     (отправлено, ctx.Err());
//   - ошибка Send → вернуть (отправлено, err) сразу;
//   - вернуть (len(items), nil) при успехе.
//
// Запуск: go run ./03_exercises/06_stream_loop
package main

import (
	"context"
	"errors"
	"fmt"
)

type sender interface {
	Send(msg string) error
}

func streamAll(ctx context.Context, items []string, s sender) (int, error) {
	// TODO: цикл с проверкой ctx.Err() и ошибок Send.
	return 0, nil
}

// fakeStream записывает сообщения; после limit отправок отменяет ctx или падает.
type fakeStream struct {
	got     []string
	limit   int
	cancel  context.CancelFunc
	sendErr error
}

func (f *fakeStream) Send(m string) error {
	if f.sendErr != nil && len(f.got) == f.limit {
		return f.sendErr
	}
	f.got = append(f.got, m)
	if f.cancel != nil && len(f.got) == f.limit {
		f.cancel()
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 06_stream_loop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	items := []string{"a", "b", "c", "d"}

	s := &fakeStream{}
	n, err := streamAll(context.Background(), items, s)
	check(err == nil && n == 4 && len(s.got) == 4, "все 4 сообщения отправлены")

	ctx, cancel := context.WithCancel(context.Background())
	s = &fakeStream{limit: 2, cancel: cancel}
	n, err = streamAll(ctx, items, s)
	check(errors.Is(err, context.Canceled) && n == 2, "клиент ушёл после 2 → остановка, Canceled")

	broken := errors.New("transport closing")
	s = &fakeStream{limit: 1, sendErr: broken}
	n, err = streamAll(context.Background(), items, s)
	check(errors.Is(err, broken) && n == 1, "ошибка Send прерывает стрим")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
