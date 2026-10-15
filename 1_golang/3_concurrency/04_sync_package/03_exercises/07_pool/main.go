// Упражнение 07: переиспользование буферов через sync.Pool.
//
// Задача: render(name) возвращает строку "hello, <name>!", собирая её во
// временном bytes.Buffer, взятом из bufPool:
//
//	buf := bufPool.Get().(*bytes.Buffer)
//	buf.Reset()               // буфер мог остаться «грязным»
//	... запись ...
//	s := buf.String()
//	bufPool.Put(buf)          // вернуть в пул
//
// Без Reset повторные вызовы получат хвост предыдущего результата.
//
// Запуск: go run ./03_exercises/07_pool/main.go
package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

// render собирает приветствие во временном буфере из пула.
func render(name string) string {
	// TODO: Get из пула, Reset, WriteString, String, Put обратно.
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_pool ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(render("gopher") == "hello, gopher!", fmt.Sprintf("render(gopher) = %q", render("gopher")))
	check(render("pool") == "hello, pool!", fmt.Sprintf("повторный вызов: %q — забыли Reset?", render("pool")))

	var wg sync.WaitGroup
	var wrong atomic.Int64
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprint("u", i)
			if render(name) != "hello, "+name+"!" {
				wrong.Add(1)
			}
		}(i)
	}
	wg.Wait()
	check(wrong.Load() == 0, fmt.Sprintf("%d конкурентных вызовов вернули неверную строку", wrong.Load()))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
