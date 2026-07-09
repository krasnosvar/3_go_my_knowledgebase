package main

import (
	"gorm.io/gorm"
)

// CreateUser — вставка нового пользователя
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUser — выборка пользователя по ID
func GetUser(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error // SELECT * FROM users WHERE id=?
	return &user, err
}

// UpdateUser — обновление данных
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error // UPDATE users SET ...
}

// DeleteUser — удаление
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error // DELETE FROM users WHERE id=?
}
