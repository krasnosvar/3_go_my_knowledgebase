// Упражнение 03: GOGC и цель следующего GC (Решение).
package main

import (
	"fmt"
	"runtime/debug"
)

// nextGCTarget возвращает размер кучи, при котором стартует следующий GC.
func nextGCTarget(liveHeap uint64, gogc int) uint64 {
	if gogc < 0 {
		return 0 // GOGC=off: по проценту GC не запускается
	}
	return liveHeap + liveHeap*uint64(gogc)/100
}

// withGCPercent выполняет f с заданным GOGC и восстанавливает прежний.
func withGCPercent(p int, f func()) {
	old := debug.SetGCPercent(p)
	defer debug.SetGCPercent(old)
	f()
}

// currentGCPercent читает текущее значение GOGC (установить и вернуть обратно).
func currentGCPercent() int {
	v := debug.SetGCPercent(100)
	debug.SetGCPercent(v)
	return v
}

func main() {
	fmt.Println("--- Проверка 03_gc_percent (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(nextGCTarget(100, 100) == 200, fmt.Sprintf("live=100, GOGC=100: ожидалось 200, получено %d", nextGCTarget(100, 100)))
	check(nextGCTarget(100, 50) == 150, fmt.Sprintf("live=100, GOGC=50: ожидалось 150, получено %d", nextGCTarget(100, 50)))
	check(nextGCTarget(64<<20, 200) == 192<<20, "live=64MB, GOGC=200: ожидалось 192MB")
	check(nextGCTarget(100, -1) == 0, "GOGC=off: ожидалось 0")

	orig := currentGCPercent()
	var inside int
	withGCPercent(50, func() { inside = currentGCPercent() })
	check(inside == 50, fmt.Sprintf("внутри f GOGC должен быть 50, получено %d", inside))
	check(currentGCPercent() == orig, fmt.Sprintf("после f GOGC должен вернуться к %d", orig))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
