package main

import (
	"fmt"
)

// runSafely запускает функцию task. Если task вызывает панику,
// runSafely должна перехватить её через recover() и вернуть её описание в виде error.
// Если паники не произошло, вернуть nil.
func runSafely(task func()) (err error) {
	// TODO: Сейчас defer + recover() только глушат панику, а err остаётся nil.
	// Доработайте: если recover() вернул значение, запишите его описание
	// в именованную переменную err (например, через fmt.Errorf("%v", r)).
	defer func() {
		_ = recover()
	}()

	task()
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_error_wrapper ---")

	// Тест 1: Функция работает штатно
	err1 := runSafely(func() {
		// Ничего не делаем
	})

	// Тест 2: Функция вызывает панику
	err2 := runSafely(func() {
		panic("что-то пошло не так")
	})

	t1 := err1 == nil
	t2 := err2 != nil && err2.Error() == "что-то пошло не так"

	if t1 && t2 {
		fmt.Println("Тест runSafely: PASS")
	} else {
		fmt.Printf("Тест runSafely: FAIL (err1=%v, err2=%v)\n", err1, err2)
	}
}
