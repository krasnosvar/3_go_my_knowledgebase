// Упражнение 11: сборка пагинированного ответа (offset/limit) (Решение).
package main

import (
	"fmt"
	"reflect"
)

type Page struct {
	Items   []int
	Limit   int
	Offset  int
	Total   int
	HasNext bool
}

func paginate(items []int, limit, offset int) Page {
	n := len(items)

	start := offset
	if start > n {
		start = n
	}
	end := start + limit
	if end > n {
		end = n
	}

	// Копируем окно в НЕ-nil срез, чтобы пустая страница была []int{}.
	window := make([]int, 0, end-start)
	window = append(window, items[start:end]...)

	return Page{
		Items:   window,
		Limit:   limit,
		Offset:  offset,
		Total:   n,
		HasNext: offset+limit < n,
	}
}

func main() {
	fmt.Println("--- Проверка 11_paginate (Решение) ---")

	data := []int{0, 1, 2, 3, 4, 5, 6}
	allOK := true

	check := func(name string, got, want Page) {
		if !reflect.DeepEqual(got, want) {
			allOK = false
			fmt.Printf("FAIL: %s\n  got : %+v\n  want: %+v\n", name, got, want)
		}
	}

	check("limit=3 offset=0",
		paginate(data, 3, 0),
		Page{Items: []int{0, 1, 2}, Limit: 3, Offset: 0, Total: 7, HasNext: true})

	check("limit=3 offset=3",
		paginate(data, 3, 3),
		Page{Items: []int{3, 4, 5}, Limit: 3, Offset: 3, Total: 7, HasNext: true})

	check("limit=3 offset=6",
		paginate(data, 3, 6),
		Page{Items: []int{6}, Limit: 3, Offset: 6, Total: 7, HasNext: false})

	got := paginate(data, 3, 100)
	want := Page{Items: []int{}, Limit: 3, Offset: 100, Total: 7, HasNext: false}
	check("offset за пределами", got, want)
	if got.Items == nil {
		allOK = false
		fmt.Println("FAIL: Items должен быть []int{} (не nil) при offset за пределами")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
