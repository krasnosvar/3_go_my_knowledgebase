package main

import (
	"fmt"
)

// Атрибуты товаров в виде битовых масок
const (
	AttrIsPhysical = 1 << 0 // 1 (0001) - Физический товар (требует доставки)
	AttrIsPremium  = 1 << 1 // 2 (0010) - Премиум товар
	AttrIsDiscount = 1 << 2 // 4 (0100) - Товар со скидкой
)

type Product struct {
	Name       string
	Attributes int // Битовая маска атрибутов
}

func main() {
	products := []Product{
		{Name: "Электронная книга Go", Attributes: AttrIsPremium | AttrIsDiscount},
		{Name: "Игровая приставка (физическая)", Attributes: AttrIsPhysical | AttrIsPremium},
		{Name: "Простой кабель", Attributes: AttrIsPhysical},
	}

	fmt.Println("--- Бэкенд: Фильтрация товаров по битовым атрибутам ---")

	for _, p := range products {
		fmt.Printf("Товар: %s\n", p.Name)

		// Проверка: Нужна ли доставка (физический товар)
		requiresShipping := (p.Attributes & AttrIsPhysical) != 0

		// Проверка: Является ли премиум товаром на скидке
		isPremiumDiscount := (p.Attributes & (AttrIsPremium | AttrIsDiscount)) == (AttrIsPremium | AttrIsDiscount)

		fmt.Printf("  -> Требует доставки: %t\n", requiresShipping)
		fmt.Printf("  -> Премиум на скидке: %t\n", isPremiumDiscount)
	}
}
