package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func main() {
	products := []Product{
		{ID: 1, Title: "Keyboard", Price: 80},
		{ID: 2, Title: "Mouse", Price: 35},
		{ID: 3, Title: "Monitor", Price: 320},
	}

	// 1. Сортировка с помощью slices.SortFunc (Go 1.21+)
	slices.SortFunc(products, func(a, b Product) int {
		return cmp.Compare(a.Price, b.Price)
	})

	fmt.Println("Отсортированные товары:")
	for _, p := range products {
		fmt.Printf("  - %s: $%d\n", p.Title, p.Price)
	}

	// 2. Бинарный поиск с помощью slices.BinarySearchFunc (только по отсортированному слайсу!)
	targetPrice := 80
	idx, found := slices.BinarySearchFunc(products, targetPrice, func(p Product, price int) int {
		return cmp.Compare(p.Price, price)
	})

	if found {
		fmt.Printf("\nТовар с ценой $%d найден на позиции %d: %s\n", targetPrice, idx, products[idx].Title)
	} else {
		fmt.Printf("\nТовар с ценой $%d не найден\n", targetPrice)
	}
}
