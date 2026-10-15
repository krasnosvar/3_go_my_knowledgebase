// Упражнение 04: keyset (cursor) пагинация.
//
// Задача: реализуйте (r *repo) ListAfter(afterID, limit int) (page []int, next int):
//   - r.ids отсортированы по возрастанию;
//   - вернуть до limit ID строго больше afterID (аналог WHERE id > $1 ORDER BY id LIMIT $2);
//   - next — последний ID страницы, если страница полная (len == limit), иначе 0
//     (0 означает «дальше данных нет»).
//
// Обход курсором от afterID=0 должен пройти все ID ровно по разу.
//
// Запуск: go run ./03_exercises/04_keyset_pagination
package main

import "fmt"

type repo struct{ ids []int }

func (r *repo) ListAfter(afterID, limit int) (page []int, next int) {
	// TODO: выберите ID > afterID, не больше limit; вычислите курсор next.
	return nil, 0
}

func main() {
	fmt.Println("--- Проверка 04_keyset_pagination ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	r := &repo{ids: []int{3, 7, 8, 15, 20}}
	p, next := r.ListAfter(0, 2)
	check(fmt.Sprint(p) == "[3 7]" && next == 7, "первая страница [3 7], next=7")
	p, next = r.ListAfter(15, 2)
	check(fmt.Sprint(p) == "[20]" && next == 0, "последняя неполная страница [20], next=0")

	var all []int
	cursor, pages := 0, 0
	for pages < 10 {
		pages++
		p, next := r.ListAfter(cursor, 2)
		all = append(all, p...)
		if next == 0 {
			break
		}
		cursor = next
	}
	check(fmt.Sprint(all) == "[3 7 8 15 20]", "обход курсором проходит все ID по разу")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
