// Упражнение 02: сборка пагинированного ответа (offset/limit).
//
// Задача: реализовать paginate — вырезать из среза items окно
// [offset, offset+limit) и собрать структуру Page с метаданными.
//
// Правила:
//   - если offset >= len(items) -> Items пустой ([]int{}), но не nil;
//   - окно не должно выходить за границы среза;
//   - Total = len(items); Limit и Offset — как переданы;
//   - HasNext = true, если после окна ещё остались элементы
//     (offset+limit < len(items)).
//
// Запуск: go run ./03_exercises/02_paginate/main.go
package main

import (
	"fmt"
	"reflect"
)

// Page — пагинированный ответ.
type Page struct {
	Items   []int
	Limit   int
	Offset  int
	Total   int
	HasNext bool
}

// paginate — реализуйте согласно условию.
func paginate(items []int, limit, offset int) Page {
	// TODO:
	//   1) start := offset; если start > len(items) -> start = len(items);
	//   2) end := start + limit; если end > len(items) -> end = len(items);
	//   3) window := items[start:end]; скопируйте в НЕ-nil срез (make + append),
	//      чтобы пустая страница была []int{}, а не nil;
	//   4) HasNext := offset+limit < len(items).
	return Page{}
}

// var _ удерживает импорт reflect используемым, пока TODO не сделан.
var _ = reflect.DeepEqual

func main() {
	fmt.Println("--- Проверка 02_paginate ---")

	data := []int{0, 1, 2, 3, 4, 5, 6} // 7 элементов
	allOK := true

	check := func(name string, got, want Page) {
		if !reflect.DeepEqual(got, want) {
			allOK = false
			fmt.Printf("FAIL: %s\n  got : %+v\n  want: %+v\n", name, got, want)
		}
	}

	// Первая страница.
	check("limit=3 offset=0",
		paginate(data, 3, 0),
		Page{Items: []int{0, 1, 2}, Limit: 3, Offset: 0, Total: 7, HasNext: true})

	// Средняя страница.
	check("limit=3 offset=3",
		paginate(data, 3, 3),
		Page{Items: []int{3, 4, 5}, Limit: 3, Offset: 3, Total: 7, HasNext: true})

	// Последняя (неполная) страница.
	check("limit=3 offset=6",
		paginate(data, 3, 6),
		Page{Items: []int{6}, Limit: 3, Offset: 6, Total: 7, HasNext: false})

	// Offset за пределами -> пустая (не nil) страница.
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
