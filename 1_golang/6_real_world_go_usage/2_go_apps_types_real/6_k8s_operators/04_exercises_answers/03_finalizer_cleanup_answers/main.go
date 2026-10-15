// Упражнение 03: finalizer и очистка внешних ресурсов (Решение).
package main

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

const finalizerName = "example.com/bucket-cleanup"

type object struct {
	Name              string
	Finalizers        []string
	DeletionTimestamp *time.Time
}

func reconcile(o *object, cleanup func(name string) error) error {
	has := slices.Contains(o.Finalizers, finalizerName)
	if o.DeletionTimestamp == nil {
		if !has {
			o.Finalizers = append(o.Finalizers, finalizerName)
		}
		return nil
	}
	if !has {
		return nil
	}
	if err := cleanup(o.Name); err != nil {
		return err
	}
	o.Finalizers = slices.DeleteFunc(o.Finalizers, func(f string) bool { return f == finalizerName })
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_finalizer_cleanup (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var cleaned []string
	ok := func(n string) error { cleaned = append(cleaned, n); return nil }
	fail := func(string) error { return errors.New("s3 unavailable") }

	o := &object{Name: "b1", Finalizers: []string{"other/keep"}}
	_ = reconcile(o, ok)
	_ = reconcile(o, ok)
	check(fmt.Sprint(o.Finalizers) == "[other/keep "+finalizerName+"]", "финализатор добавлен один раз")

	now := time.Now()
	o.DeletionTimestamp = &now
	err := reconcile(o, fail)
	check(err != nil && slices.Contains(o.Finalizers, finalizerName), "ошибка cleanup → финализатор остаётся")
	err = reconcile(o, ok)
	check(err == nil && fmt.Sprint(o.Finalizers) == "[other/keep]", "успешный cleanup → снят только наш финализатор")
	check(fmt.Sprint(cleaned) == "[b1]", "cleanup вызван для b1")
	_ = reconcile(o, ok)
	check(len(cleaned) == 1, "без финализатора cleanup не повторяется")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
