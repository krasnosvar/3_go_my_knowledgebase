package main

import (
	"fmt"
)

// UserRole представляет битовую маску ролей/прав пользователя
type UserRole int

const (
	RoleRead   UserRole                            = 1 << iota // 1 << 0 == 0001 (1)
	RoleWrite                                                  // 1 << 1 == 0010 (2)
	RoleManage                                                 // 1 << 2 == 0100 (4)
	RoleAdmin  = RoleRead | RoleWrite | RoleManage             // 0111 (7)
)

// HasRole проверяет наличие конкретной роли у пользователя через битовое И
func HasRole(userRoles UserRole, check UserRole) bool {
	return (userRoles & check) == check
}

func main() {
	// Обычный менеджер: может читать и управлять, но не писать напрямую в файлы
	var manager UserRole = RoleRead | RoleManage

	// Суперпользователь
	var superUser UserRole = RoleAdmin

	fmt.Println("--- Проверка битовых масок прав доступа ---")
	fmt.Printf("Менеджер может читать: %t\n", HasRole(manager, RoleRead))
	fmt.Printf("Менеджер может писать: %t\n", HasRole(manager, RoleWrite))
	fmt.Printf("Менеджер может управлять: %t\n", HasRole(manager, RoleManage))

	fmt.Println()
	fmt.Printf("Админ может читать: %t\n", HasRole(superUser, RoleRead))
	fmt.Printf("Админ может писать: %t\n", HasRole(superUser, RoleWrite))
	fmt.Printf("Админ может управлять: %t\n", HasRole(superUser, RoleManage))
}
