package main

import (
	"fmt"
)

// Описание битовых масок прав доступа
const (
	PermissionRead    = 1 << 0 // 1  (0001)
	PermissionWrite   = 1 << 1 // 2  (0010)
	PermissionExecute = 1 << 2 // 4  (0100)
)

func main() {
	fmt.Println("--- CLI: Управление битовыми флагами прав доступа ---")

	// Начинаем с пустых прав
	var userPerms = 0

	// 1. Добавляем права на чтение и выполнение с помощью побитового ИЛИ (|)
	userPerms |= PermissionRead
	userPerms |= PermissionExecute
	fmt.Printf("Права пользователя (двоичный вид): %04b (код: %d)\n", userPerms, userPerms)

	// 2. Проверяем наличие прав с помощью побитового И (&)
	canRead := (userPerms & PermissionRead) != 0
	canWrite := (userPerms & PermissionWrite) != 0
	fmt.Printf("Пользователь может читать: %t\n", canRead)
	fmt.Printf("Пользователь может писать: %t\n", canWrite)

	// 3. Отзываем права на выполнение с помощью побитового И НЕ (&^)
	fmt.Println("\nОтзываем права на выполнение...")
	userPerms &^= PermissionExecute
	fmt.Printf("Права после отзыва: %04b\n", userPerms)

	canExecute := (userPerms & PermissionExecute) != 0
	fmt.Printf("Пользователь может выполнять: %t\n", canExecute)
}
