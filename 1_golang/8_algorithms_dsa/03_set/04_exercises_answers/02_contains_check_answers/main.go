// Упражнение 02: множество строк: добавление и проверка (Решение).
package main

import "fmt"

// StringSet — множество строк.
type StringSet map[string]struct{}

// newStringSet создаёт множество из перечисленных элементов.
func newStringSet(items ...string) StringSet {
	s := make(StringSet, len(items))
	for _, it := range items {
		s.Add(it)
	}
	return s
}

// Add добавляет элемент в множество.
func (s StringSet) Add(x string) {
	s[x] = struct{}{}
}

// Has сообщает, есть ли элемент в множестве.
func (s StringSet) Has(x string) bool {
	_, ok := s[x]
	return ok
}

func main() {
	fmt.Println("--- Проверка 02_contains_check (Решение) ---")

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
