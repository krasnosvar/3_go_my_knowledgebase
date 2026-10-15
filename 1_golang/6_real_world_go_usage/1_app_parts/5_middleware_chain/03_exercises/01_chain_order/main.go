// Упражнение 01: сборка цепочки middleware.
//
// Задача: реализуйте chain(h http.Handler, mws ...middleware) http.Handler так,
// чтобы ПЕРВЫЙ middleware в списке был ВНЕШНИМ (выполнялся первым на входе
// и последним на выходе): chain(h, a, b) == a(b(h)).
//
// Запуск: go run ./03_exercises/01_chain_order
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type middleware func(http.Handler) http.Handler

func chain(h http.Handler, mws ...middleware) http.Handler {
	// TODO: оберните h в обратном порядке, чтобы mws[0] оказался снаружи.
	return h
}

func main() {
	fmt.Println("--- Проверка 01_chain_order ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var trace []string
	mark := func(name string) middleware {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				trace = append(trace, name+">")
				next.ServeHTTP(w, r)
				trace = append(trace, "<"+name)
			})
		}
	}
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { trace = append(trace, "h") })

	chain(final, mark("a"), mark("b")).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
	got := strings.Join(trace, " ")
	check(got == "a> b> h <b <a", "порядок a> b> h <b <a, получено: "+got)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
