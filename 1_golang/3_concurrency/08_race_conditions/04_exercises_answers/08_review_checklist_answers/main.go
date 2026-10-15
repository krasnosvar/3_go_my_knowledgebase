// Упражнение 08: ревью конкурентного кода (квиз) (Решение).
package main

import "fmt"

var answers = map[int]string{
	1: "b",
	2: "b",
	3: "a",
	4: "b",
	5: "a",
}

func main() {
	fmt.Println("--- Проверка 08_review_checklist (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	key := map[int]string{1: "b", 2: "b", 3: "a", 4: "b", 5: "a"}
	for q := 1; q <= 5; q++ {
		got, ok := answers[q]
		if !ok {
			check(false, fmt.Sprintf("вопрос %d: нет ответа", q))
			continue
		}
		check(got == key[q], fmt.Sprintf("вопрос %d: ответ %q неверен", q, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
