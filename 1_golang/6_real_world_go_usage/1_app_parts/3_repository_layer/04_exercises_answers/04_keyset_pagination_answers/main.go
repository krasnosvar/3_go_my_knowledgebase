// Упражнение 04: keyset (cursor) пагинация (Решение).
package main

import "fmt"

type repo struct{ ids []int }

func (r *repo) ListAfter(afterID, limit int) (page []int, next int) {
	for _, id := range r.ids {
		if id > afterID {
			page = append(page, id)
			if len(page) == limit {
				return page, id
			}
		}
	}
	return page, 0
}

func main() {
	fmt.Println("--- Проверка 04_keyset_pagination (Решение) ---")

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
