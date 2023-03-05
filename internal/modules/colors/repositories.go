package colors

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateColor(color *models.Color) (result *gorm.DB) {
	return databases.PostgresDB.Create(color)
}

func ListColors(colors *[]models.Color) (result *gorm.DB) {
	return databases.PostgresDB.Find(colors)
}

func FindColor(id int, color *models.Color) (result *gorm.DB) {
	return databases.PostgresDB.First(color, id)
}

func UpdateColor(id int, color *models.Color) (result *gorm.DB) {
	databases.PostgresDB.Model(color).Where("id = ?", id).Updates(color)
	return FindColor(id, color)
}

func DeleteColor(color *models.Color) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(color)
}
