package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func findPointName(db map[Point]string, pt Point) (string, bool) {
	name, ok := db[pt]
	if !ok {
		return "unknown", false
	}
	return name, true
}

func main() {
	fmt.Println("--- Проверка 05_maps_key_constraints (Решение) ---")

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
