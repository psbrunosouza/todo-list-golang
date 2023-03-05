package users

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) (result *gorm.DB) {
	return databases.PostgresDB.Create(user)
}

func ListUsers(users *[]models.User) (result *gorm.DB) {
	return databases.PostgresDB.Find(users)
}

func FindUser(id int, user *models.User) (result *gorm.DB) {
	return databases.PostgresDB.First(user, id)
}

func UpdateUser(id int, user *models.User) (result *gorm.DB) {
	return databases.PostgresDB.Model(user).Where("id = ?", id).Update("User", user)
}

func DeleteUser(user *models.User) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(user)
}
