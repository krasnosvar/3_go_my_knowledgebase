// Упражнение 02: множество строк: добавление и проверка.
//
// Задача: реализуйте множество строк на map[string]struct{}:
//   - newStringSet(items...) создаёт множество из элементов;
//   - метод Has(x) сообщает, есть ли элемент;
//   - метод Add(x) добавляет элемент.
//
// struct{} занимает 0 байт — идиоматичный способ хранить «только ключи».
//
// Запуск: go run ./03_exercises/02_contains_check/main.go
package main

import "fmt"

// StringSet — множество строк.
type StringSet map[string]struct{}

// newStringSet создаёт множество из перечисленных элементов.
func newStringSet(items ...string) StringSet {
	// TODO: создайте map и добавьте все items.
	return StringSet{}
}

// Add добавляет элемент в множество.
func (s StringSet) Add(x string) {
	// TODO: s[x] = struct{}{}
}

// Has сообщает, есть ли элемент в множестве.
func (s StringSet) Has(x string) bool {
	// TODO: comma-ok проверка.
	return false
}

func main() {
	fmt.Println("--- Проверка 02_contains_check ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := newStringSet("go", "rust")
	check(s.Has("go"), "Has(\"go\"): ожидалось true")
	check(!s.Has("java"), "Has(\"java\"): ожидалось false")

	s.Add("java")
	check(s.Has("java"), "после Add(\"java\") Has должен вернуть true")
	check(len(s) == 3, fmt.Sprintf("ожидалось 3 элемента, получено %d", len(s)))

	s.Add("go")
	check(len(s) == 3, "повторный Add не должен увеличивать размер")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
