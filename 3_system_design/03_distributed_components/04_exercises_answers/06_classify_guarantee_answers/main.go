// Упражнение 06: классификация гарантии доставки (Решение).
package main

import "fmt"

func classify(ackFirst, dedup bool) string {
	// ack ДО обработки -> возможна потеря -> at-most-once (дедуп не помогает).
	if ackFirst {
		return "at-most-once"
	}
	// ack ПОСЛЕ обработки -> возможны дубли. Есть дедуп -> effectively
	// exactly-once.
	if dedup {
		return "exactly-once"
	}
	return "at-least-once"
}

var _ = classify

func main() {
	fmt.Println("--- Проверка 06_classify_guarantee (Решение) ---")

	cases := []struct {
		ackFirst bool
		dedup    bool
		want     string
	}{
		{true, false, "at-most-once"},
		{true, true, "at-most-once"},
		{false, false, "at-least-once"},
		{false, true, "exactly-once"},
	}

	allOK := true
	for _, c := range cases {
		got := classify(c.ackFirst, c.dedup)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: classify(ackFirst=%v, dedup=%v) = %q, ожидалось %q\n",
				c.ackFirst, c.dedup, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
