// Упражнение 03: offset-пагинация в репозитории.
//
// Задача: реализуйте (r *repo) List(offset, limit int) []item:
//   - элементы упорядочены по ID (r.items уже отсортированы);
//   - вернуть не более limit элементов, начиная с позиции offset;
//   - offset за пределами → пустой (не nil) срез;
//   - результат — КОПИЯ: изменение возвращённого среза не должно
//     менять r.items.
//
// Запуск: go run ./03_exercises/03_offset_pagination
package main

import "fmt"

type item struct {
	ID   int
	Name string
}

type repo struct{ items []item }

func (r *repo) List(offset, limit int) []item {
	// TODO: ограничьте границы и верните копию подсреза.
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_offset_pagination ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := &repo{}
	for i := 1; i <= 5; i++ {
		r.items = append(r.items, item{ID: i, Name: fmt.Sprint("n", i)})
	}

	p := r.List(0, 2)
	check(len(p) == 2 && p[0].ID == 1 && p[1].ID == 2, "List(0,2) → [1,2]")
	p = r.List(4, 10)
	check(len(p) == 1 && p[0].ID == 5, "List(4,10) → [5]")
	p = r.List(10, 2)
	check(p != nil && len(p) == 0, "List(10,2) → пустой не-nil срез")
	p = r.List(0, 1)
	if len(p) == 1 {
		p[0].Name = "hacked"
	}
	check(r.items[0].Name == "n1", "изменение результата не трогает хранилище")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
