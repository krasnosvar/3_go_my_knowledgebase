// Упражнение 02: status conditions (Решение).
package main

import (
	"fmt"
	"time"
)

type condition struct {
	Type, Status, Reason, Message string
	LastTransitionTime            time.Time
}

func setCondition(conds []condition, c condition, now time.Time) []condition {
	for i := range conds {
		if conds[i].Type != c.Type {
			continue
		}
		if conds[i].Status != c.Status {
			conds[i].Status = c.Status
			conds[i].LastTransitionTime = now
		}
		conds[i].Reason = c.Reason
		conds[i].Message = c.Message
		return conds
	}
	c.LastTransitionTime = now
	return append(conds, c)
}

func main() {
	fmt.Println("--- Проверка 02_status_conditions (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t1, t2 := t0.Add(time.Minute), t0.Add(2*time.Minute)

	var conds []condition
	conds = setCondition(conds, condition{Type: "Ready", Status: "False", Reason: "Starting"}, t0)
	check(len(conds) == 1 && conds[0].LastTransitionTime.Equal(t0), "новое условие добавлено с t0")

	conds = setCondition(conds, condition{Type: "Ready", Status: "False", Reason: "Pulling"}, t1)
	if len(conds) == 1 {
		check(conds[0].Reason == "Pulling" && conds[0].LastTransitionTime.Equal(t0), "тот же Status → Reason обновлён, время прежнее")
	}
	conds = setCondition(conds, condition{Type: "Ready", Status: "True", Reason: "Running"}, t2)
	if len(conds) == 1 {
		check(conds[0].Status == "True" && conds[0].LastTransitionTime.Equal(t2), "смена Status → новое время перехода")
	}
	conds = setCondition(conds, condition{Type: "Degraded", Status: "False"}, t2)
	check(len(conds) == 2, "условия уникальны по Type, новый Type добавлен")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
