package main

import (
	"fmt"
	"unsafe"
)

// BadStruct имеет неоптимальный порядок полей (размер 24 байта)
type BadStruct struct {
	a bool  // 1 байт
	b int64 // 8 байт
	c bool  // 1 байт
}

// GoodStruct сгруппирована по убыванию размера полей (размер 16 байт)
type GoodStruct struct {
	b int64 // 8 байт
	a bool  // 1 байт
	c bool  // 1 байт
}

func main() {
	badSize := unsafe.Sizeof(BadStruct{})
	goodSize := unsafe.Sizeof(GoodStruct{})

	fmt.Printf("Размер BadStruct в памяти:  %d байт\n", badSize)
	fmt.Printf("Размер GoodStruct в памяти: %d байт\n", goodSize)
	fmt.Printf("Экономия памяти: %.2f%%\n", float64(badSize-goodSize)/float64(badSize)*100)
}
