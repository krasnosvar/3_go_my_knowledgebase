// Упражнение 03: определить CP или AP по описанию поведения при разделении сети.
//
// Задача: реализуйте classifyCAP(behavior) — по описанию того, как система
// ведёт себя во время сетевого разделения (partition), верните:
//   - "CP", если система жертвует ДОСТУПНОСТЬЮ ради согласованности
//     (узел, не уверенный в свежести данных, ОТКАЗЫВАЕТ в ответе / ждёт согласования);
//   - "AP", если система жертвует СОГЛАСОВАННОСТЬЮ ради доступности
//     (узел ОТВЕЧАЕТ тем, что знает, даже устаревшими данными);
//   - "unknown", если поведение не подходит ни под один случай.
//
// Вход — структура с двумя булевыми признаками (см. behavior ниже),
// решение — по ним классифицировать. Реальный "разбор текста" не нужен:
// признаки уже извлечены.
//
// Запуск: go run ./03_exercises/03_classify_cap/main.go
package main

import "fmt"

// behavior — извлечённые признаки поведения узла во время partition.
type behavior struct {
	// respondsDuringPartition: отвечает ли узел на запросы во время разделения.
	respondsDuringPartition bool
	// mayReturnStale: может ли вернуть устаревшие (несогласованные) данные.
	mayReturnStale bool
}

// classifyCAP классифицирует систему как "CP", "AP" или "unknown".
//
// Логика:
//   - отвечает во время разделения и допускает устаревшие данные -> "AP";
//   - НЕ отвечает во время разделения (жертвует доступностью ради согласованности) -> "CP";
//   - иначе (отвечает, но при этом гарантирует свежесть при разделении —
//     что невозможно по CAP) -> "unknown".
func classifyCAP(b behavior) string {
	// TODO: верните "CP" / "AP" / "unknown" по признакам b.
	return ""
}

func main() {
	fmt.Println("--- Проверка 03_classify_cap ---")

	cases := []struct {
		name string
		b    behavior
		want string
	}{
		// Dynamo/Cassandra: всегда отвечает, может отдать устаревшее -> AP.
		{"cassandra", behavior{respondsDuringPartition: true, mayReturnStale: true}, "AP"},
		// etcd/ZooKeeper (кворум): меньшинство отказывает в ответе -> CP.
		{"etcd", behavior{respondsDuringPartition: false, mayReturnStale: false}, "CP"},
		// Отказывает, но якобы отдаёт устаревшее — всё равно не отвечает -> CP.
		{"cp-variant", behavior{respondsDuringPartition: false, mayReturnStale: true}, "CP"},
		// Отвечает и гарантирует свежесть при разделении — по CAP невозможно -> unknown.
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
