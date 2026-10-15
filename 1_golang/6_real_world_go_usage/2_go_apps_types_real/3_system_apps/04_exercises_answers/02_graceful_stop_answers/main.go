// Упражнение 02: упорядоченная остановка компонентов (Решение).
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

func shutdown(ctx context.Context, comps []component) error {
	var errs []error
	for i := len(comps) - 1; i >= 0; i-- {
		c := comps[i]
		if err := ctx.Err(); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", c.name, err))
			continue
		}
		if err := c.Stop(ctx); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", c.name, err))
		}
	}
	return errors.Join(errs...)
}

func main() {
	fmt.Println("--- Проверка 02_graceful_stop (Решение) ---")

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
