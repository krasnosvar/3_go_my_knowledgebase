// Упражнение 07: переиспользование буферов через sync.Pool (Решение).
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

func render(name string) string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString("hello, ")
	buf.WriteString(name)
	buf.WriteString("!")
	s := buf.String()
	bufPool.Put(buf)
	return s
}

func main() {
	fmt.Println("--- Проверка 07_pool (Решение) ---")

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
