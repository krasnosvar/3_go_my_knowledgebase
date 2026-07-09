package main

import (
	"fmt"
	"sync"
)

type buf struct{ b []byte }

var pool = sync.Pool{New: func() any { return &buf{b: make([]byte, 0, 64)} }}

func main() {
	v := pool.Get().(*buf)
	v.b = append(v.b, 'o', 'k')
	fmt.Println(string(v.b))
	v.b = v.b[:0]
	pool.Put(v)
}
