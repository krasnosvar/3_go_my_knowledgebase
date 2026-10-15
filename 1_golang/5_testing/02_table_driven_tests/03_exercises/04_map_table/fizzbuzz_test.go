// Упражнение 04: таблица-map и t.Run.
//
// Задача: напишите TestFizzBuzz, где таблица — map[string]struct{in int;
// want string}: ключ map — имя случая. Каждый случай запускайте через
// t.Run(name, ...), чтобы можно было выполнить один:
//
//	go test -run 'TestFizzBuzz/fifteen' ./03_exercises/04_map_table/
//
// Порядок обхода map случаен — это плюс: тесты не должны зависеть от порядка.
// Случаи: 1, 3, 5, 15, 30, 7.
//
// Запуск: go test ./03_exercises/04_map_table/
package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Fatal("TODO: таблица в виде map + t.Run")
}
