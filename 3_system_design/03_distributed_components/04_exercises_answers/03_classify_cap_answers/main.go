// Упражнение 03: определить CP или AP по описанию (Решение).
package main

import "fmt"

type behavior struct {
	respondsDuringPartition bool
	mayReturnStale          bool
}

func classifyCAP(b behavior) string {
	// Жертвует доступностью (не отвечает при разделении) -> CP.
	if !b.respondsDuringPartition {
		return "CP"
	}
	// Отвечает и допускает устаревшие данные -> AP.
	if b.mayReturnStale {
		return "AP"
	}
	// Отвечает и гарантирует свежесть при разделении — по CAP невозможно.
	return "unknown"
}

func main() {
	fmt.Println("--- Проверка 03_classify_cap (Решение) ---")

	cases := []struct {
		name string
		b    behavior
		want string
	}{
		{"cassandra", behavior{respondsDuringPartition: true, mayReturnStale: true}, "AP"},
		{"etcd", behavior{respondsDuringPartition: false, mayReturnStale: false}, "CP"},
		{"cp-variant", behavior{respondsDuringPartition: false, mayReturnStale: true}, "CP"},
		{"impossible", behavior{respondsDuringPartition: true, mayReturnStale: false}, "unknown"},
	}

	allOK := true
	for _, c := range cases {
		got := classifyCAP(c.b)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: classifyCAP(%s) = %q, ожидалось %q\n", c.name, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
