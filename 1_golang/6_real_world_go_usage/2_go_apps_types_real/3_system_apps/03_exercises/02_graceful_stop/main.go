// Упражнение 02: упорядоченная остановка компонентов.
//
// Компоненты запускались в порядке db → cache → http; останавливать их
// нужно в ОБРАТНОМ порядке (сначала перестать принимать запросы).
//
// Задача: реализуйте shutdown(ctx, comps []component) error:
//   - вызвать Stop(ctx) у компонентов с конца к началу;
//   - ошибка одного компонента не мешает остановить остальные;
//     все ошибки собрать через errors.Join;
//   - если ctx уже истёк перед очередным компонентом — не вызывать его Stop,
//     а добавить ошибку fmt.Errorf("%s: %w", name, ctx.Err()).
//
// Запуск: go run ./03_exercises/02_graceful_stop
package main

import (
	"context"
	"errors"
	"fmt"
)

type component struct {
	name string
	Stop func(ctx context.Context) error
}

var _ = errors.Join

func shutdown(ctx context.Context, comps []component) error {
	// TODO: обратный порядок, errors.Join, проверка ctx.Err() перед каждым.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_graceful_stop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var order []string
	errCache := errors.New("flush failed")
	mk := func(name string, err error) component {
		return component{name: name, Stop: func(context.Context) error { order = append(order, name); return err }}
	}
	comps := []component{mk("db", nil), mk("cache", errCache), mk("http", nil)}

	err := shutdown(context.Background(), comps)
	check(fmt.Sprint(order) == "[http cache db]", "обратный порядок остановки: "+fmt.Sprint(order))
	check(errors.Is(err, errCache), "ошибка cache возвращена, db всё равно остановлен")

	order = nil
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = shutdown(ctx, comps)
	check(len(order) == 0 && errors.Is(err, context.Canceled), "истёкший ctx → Stop не вызывается, ошибка Canceled")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
