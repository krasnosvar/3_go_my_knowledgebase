// Упражнение 04: sync.Pool для буферов (Решение).
package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func getBuf() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func putBuf(b *bytes.Buffer) {
	b.Reset() // без Reset следующий Get увидит старое содержимое
	bufPool.Put(b)
}

func render(name string) string {
	b := getBuf()
	defer putBuf(b)
	b.WriteString("Hello, ")
	b.WriteString(name)
	b.WriteString("!")
	return b.String() // String() копирует байты — буфер можно отдавать в пул
}

func main() {
	fmt.Println("--- Проверка 04_pool_usage (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(render("Go") == "Hello, Go!", fmt.Sprintf("render(\"Go\"): ожидалось \"Hello, Go!\", получено %q", render("Go")))
	check(render("Pool") == "Hello, Pool!", "повторный render не должен содержать хвост прошлого вызова")

	dirty := getBuf()
	dirty.WriteString("secret")
	putBuf(dirty)
	again := getBuf()
	check(again.Len() == 0, fmt.Sprintf("буфер из пула должен быть пустым, а в нём %q", again.String()))
	putBuf(again)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
