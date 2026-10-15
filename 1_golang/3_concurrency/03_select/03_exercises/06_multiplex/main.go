// Упражнение 06: мультиплексирование N источников.
//
// Задача: multiplex(srcs...) объединяет произвольное число каналов-источников
// в один выходной. Выход закрывается, когда закрыты все источники.
//
// Подсказка: по горутине на источник + sync.WaitGroup + отдельная горутина,
// закрывающая выход после wg.Wait(). Порядок значений не важен.
//
// Запуск: go run ./03_exercises/06_multiplex/main.go
package main

import (
	"fmt"
	"sort"
	"time"
)

// multiplex сливает все srcs в один канал.
func multiplex(srcs ...<-chan string) <-chan string {
	// TODO: горутина на каждый источник, close(out) после wg.Wait().
	return nil
}

// --- вспомогательные функции проверки (не меняйте) ---

// collectStrings читает канал до закрытия; ok=false, если не закрылся за d.
func collectStrings(ch <-chan string, d time.Duration) ([]string, bool) {
	res := []string{}
	deadline := time.After(d)
	for {
		select {
		case v, open := <-ch:
			if !open {
				return res, true
			}
			res = append(res, v)
		case <-deadline:
			return res, false
		}
	}
}

func main() {
	fmt.Println("--- Проверка 06_multiplex ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	mk := func(vals ...string) <-chan string {
		ch := make(chan string, len(vals))
		for _, v := range vals {
			ch <- v
		}
		close(ch)
		return ch
	}

	got, closed := collectStrings(multiplex(mk("a1", "a2"), mk("b1"), mk("c1", "c2", "c3")), 300*time.Millisecond)
	sort.Strings(got)
	check(closed, "выходной канал не закрылся")
	check(fmt.Sprint(got) == "[a1 a2 b1 c1 c2 c3]", fmt.Sprintf("получено %v, ожидалось все 6 значений", got))

	got, closed = collectStrings(multiplex(), 300*time.Millisecond)
	check(closed && len(got) == 0, "без источников выход должен сразу закрыться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
