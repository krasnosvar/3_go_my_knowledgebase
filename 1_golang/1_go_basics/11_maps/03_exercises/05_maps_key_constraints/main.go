package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

// findPointName ищет название точки по ее координатам Point в мапе.
// Если точки нет в базе, верните "unknown", false.
func findPointName(db map[Point]string, pt Point) (string, bool) {
	// TODO: Реализуйте поиск по ключу-структуре.
	return "", false
}

func main() {
	fmt.Println("--- Проверка 05_maps_key_constraints ---")

	db := map[Point]string{
		{X: 10, Y: 20}: "Start",
		{X: 30, Y: 40}: "Finish",
	}

	name1, ok1 := findPointName(db, Point{X: 10, Y: 20})
	name2, ok2 := findPointName(db, Point{X: 99, Y: 99})

	t1 := ok1 && name1 == "Start"
	t2 := !ok2 && name2 == "unknown"

	if t1 && t2 {
		fmt.Println("Тест Point keys: PASS")
	} else {
		fmt.Printf("Тест Point keys: FAIL (name1=%q, ok1=%t, name2=%q)\n", name1, ok1, name2)
	}
}
