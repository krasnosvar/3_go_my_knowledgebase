// Упражнение 02: почему переменная убегает в кучу (Решение).
package main

import "fmt"

const (
	reasonReturnPointer = "return_pointer" // наружу возвращается указатель на локальную переменную
	reasonGlobal        = "global"         // значение сохраняется в глобальную переменную
	reasonClosure       = "closure"        // переменную захватило замыкание, которое переживает функцию
	reasonInterface     = "interface"      // значение упаковано в интерфейс, который уходит наружу
	reasonNone          = "none"           // не убегает: остаётся на стеке
)

type user struct{ name string }

var registry []*user

var lastValue any

// newUser: переменная u.
func newUser(name string) *user {
	u := user{name: name}
	return &u
}

// register: переменная u.
func register(name string) {
	u := &user{name: name}
	registry = append(registry, u)
}

// counter: переменная n.
func counter() func() int {
	n := 0
	return func() int { n++; return n }
}

// remember: значение v.
func remember(v int) {
	lastValue = v
}

// localSum: массив arr.
func localSum() int {
	arr := [4]int{1, 2, 3, 4}
	return arr[0] + arr[1] + arr[2] + arr[3]
}

var reasons = map[string]string{
	"newUser":  reasonReturnPointer,
	"register": reasonGlobal,
	"counter":  reasonClosure,
	"remember": reasonInterface,
	"localSum": reasonNone,
}

func main() {
	fmt.Println("--- Проверка 02_escape_reason (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	_ = newUser("a")
	register("b")
	_ = counter()()
	remember(1000)
	_ = localSum()

	expected := map[string]string{
		"newUser":  reasonReturnPointer,
		"register": reasonGlobal,
		"counter":  reasonClosure,
		"remember": reasonInterface,
		"localSum": reasonNone,
	}
	for _, name := range []string{"newUser", "register", "counter", "remember", "localSum"} {
		check(reasons[name] == expected[name], fmt.Sprintf("%s: ожидалось %q, указано %q", name, expected[name], reasons[name]))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
