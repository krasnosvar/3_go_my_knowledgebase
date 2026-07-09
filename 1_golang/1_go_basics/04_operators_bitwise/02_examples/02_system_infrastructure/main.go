package main

import (
	"fmt"
)

// Описание битовых флагов системных предупреждений
const (
	FlagCPUOverload = 1 << iota // 0001
	FlagOutOfMemory             // 0010
	FlagDiskFull                // 0100
)

func main() {
	fmt.Println("--- Системный мониторинг: Флаги состояний ---")

	// Регистр состояния сервера (базовый тип uint8)
	var statusRegistry uint8 = 0

	// 1. Возникли проблемы с памятью и процессором
	statusRegistry |= FlagOutOfMemory | FlagCPUOverload
	fmt.Printf("Регистр состояния предупреждений: %04b\n", statusRegistry)

	// 2. Функция парсинга и логирования
	logAlerts(statusRegistry)

	// 3. Диск очищен, ОЗУ освободилось, процессор всё ещё перегружен
	// Сбрасываем флаг памяти и включаем/сбрасываем флаг диска
	statusRegistry &^= FlagOutOfMemory
	fmt.Println("\nОчистили память. Текущее состояние:")
	logAlerts(statusRegistry)
}

func logAlerts(status uint8) {
	if status == 0 {
		fmt.Println("Система работает стабильно. Предупреждений нет.")
		return
	}

	if (status & FlagCPUOverload) != 0 {
		fmt.Println("[ALERT] Превышен лимит загрузки CPU!")
	}
	if (status & FlagOutOfMemory) != 0 {
		fmt.Println("[CRITICAL] Недостаточно оперативной памяти (OOM)!")
	}
	if (status & FlagDiskFull) != 0 {
		fmt.Println("[WARNING] Диск заполнен более чем на 95%!")
	}
}
