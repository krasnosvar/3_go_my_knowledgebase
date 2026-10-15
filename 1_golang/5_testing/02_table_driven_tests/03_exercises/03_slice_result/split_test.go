// Упражнение 03: сравнение слайсов в таблице.
//
// Задача: напишите табличный TestSplitTrim с полем want []string.
// Слайсы нельзя сравнить через ==: используйте slices.Equal (nil и пустой
// слайс он считает равными) или reflect.DeepEqual (не считает!).
// Случаи: "a,b,c"; " a , b "; "a,,b"; ""; ",,,".
//
// Запуск: go test ./03_exercises/03_slice_result/
package split

import "testing"

func TestSplitTrim(t *testing.T) {
	t.Fatal("TODO: табличный тест со слайсом в ожидании")
}
