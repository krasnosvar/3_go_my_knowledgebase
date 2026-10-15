// Упражнение 02: счётчик с лейблами и вывод в формате Prometheus.
//
// Задача:
//   - (c *counter) Inc(method, code string) — потокобезопасно (mutex)
//     увеличить значение для пары лейблов;
//   - (c *counter) Render() string — строки вида
//     http_requests_total{code="200",method="GET"} 3
//     по одной на пару, отсортированные лексикографически, каждая с "\n".
//
// Тест вызывает Inc из 50 горутин.
//
// Запуск: go run ./03_exercises/02_counter_exposition
package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type counter struct {
	mu   sync.Mutex
	vals map[string]int
}

var _ = sort.Strings
var _ = strings.Builder{}

func (c *counter) Inc(method, code string) {
	// TODO: под мьютексом увеличьте c.vals[ключ из лейблов].
}

func (c *counter) Render() string {
	// TODO: отсортируйте ключи и соберите строки "<ключ> <значение>\n".
	return ""
}

func main() {
	fmt.Println("--- Проверка 02_counter_exposition ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := &counter{vals: map[string]int{}}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%10 == 0 {
				c.Inc("POST", "500")
			} else {
				c.Inc("GET", "200")
			}
		}(i)
	}
	wg.Wait()

	want := "http_requests_total{code=\"200\",method=\"GET\"} 45\n" +
		"http_requests_total{code=\"500\",method=\"POST\"} 5\n"
	got := c.Render()
	check(got == want, "Render:\n"+got)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
