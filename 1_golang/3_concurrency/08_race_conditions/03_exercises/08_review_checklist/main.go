// Упражнение 08: ревью конкурентного кода (квиз).
//
// Задача: ответьте на вопросы ревью, заполнив карту answers буквами
// вариантов ("a", "b" или "c"). Подсказки — в теории темы.
//
//  1. В map пишут из нескольких горутин без синхронизации. Что произойдёт?
//     a) ничего страшного, map в Go потокобезопасна
//     b) data race; runtime может упасть с "concurrent map writes"
//     c) компилятор не соберёт такой код
//  2. Как надёжнее всего обнаружить data race?
//     a) go vet
//     b) запуск тестов/программы с флагом -race
//     c) добавить time.Sleep и посмотреть на вывод
//  3. Две горутины захватывают мьютексы A и B в разном порядке. Риск и лечение?
//     a) deadlock; единый порядок захвата
//     b) livelock; увеличить GOMAXPROCS
//     c) никакого риска
//  4. Структура с sync.Mutex передаётся в функцию по значению. Что не так?
//     a) всё в порядке
//     b) мьютекс копируется — защита не работает (go vet: copylocks); нужен указатель
//     c) это ошибка компиляции
//  5. Горутина пишет в канал, который никто никогда не прочитает. Что это?
//     a) утечка горутины: она навсегда заблокирована
//     b) runtime её автоматически завершит
//     c) значение будет отброшено, горутина продолжит
//
// Запуск: go run ./03_exercises/08_review_checklist/main.go
package main

import "fmt"

// answers: номер вопроса → буква ответа.
var answers = map[int]string{
	// TODO: 1: "?", 2: "?", ...
}

func main() {
	fmt.Println("--- Проверка 08_review_checklist ---")

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
