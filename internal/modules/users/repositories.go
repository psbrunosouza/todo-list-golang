package users

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

func CreateUser(user *entities.User) (result *gorm.DB) {
	return databases.PostgresDB.Create(user)
}

func ListUsers(users *[]entities.User) (result *gorm.DB) {
	return databases.PostgresDB.Find(users)
}

func FindUser(id int, user *entities.User) (result *gorm.DB) {
	return databases.PostgresDB.First(user, id)
}

func UpdateUser(id int, user *entities.User) (result *gorm.DB) {
	databases.PostgresDB.Model(user).Where("id = ?", id).Updates(user)
	return FindUser(id, user)
}

func DeleteUser(user *entities.User) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(user)
}
