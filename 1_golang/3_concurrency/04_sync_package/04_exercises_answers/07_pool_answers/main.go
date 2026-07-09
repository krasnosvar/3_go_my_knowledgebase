package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{New: func() any { return new(bytes.Buffer) }}
	b := pool.Get().(*bytes.Buffer)
	b.WriteString("ok")
	fmt.Println(b.String())
	b.Reset()
	pool.Put(b)
}
