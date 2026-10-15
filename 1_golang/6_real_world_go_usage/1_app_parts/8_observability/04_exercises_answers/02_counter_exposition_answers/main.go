// Упражнение 02: счётчик с лейблами (Решение).
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

func (c *counter) Inc(method, code string) {
	key := fmt.Sprintf(`http_requests_total{code=%q,method=%q}`, code, method)
	c.mu.Lock()
	c.vals[key]++
	c.mu.Unlock()
}

func (c *counter) Render() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	keys := make([]string, 0, len(c.vals))
	for k := range c.vals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&b, "%s %d\n", k, c.vals[k])
	}
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 02_counter_exposition (Решение) ---")

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
