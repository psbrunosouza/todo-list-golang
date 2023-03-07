package iterations

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateIteration(iteration *models.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Create(iteration)
}

func ListIterations(iteration *[]models.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Find(iteration)
}

func FindIteration(id int, iteration *models.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.First(iteration, id)
}

func UpdateIteration(id int, iteration *models.Iteration) (result *gorm.DB) {
	databases.PostgresDB.Model(iteration).Where("id = ?", id).Updates(iteration)
	return FindIteration(id, iteration)
}

func DeleteIteration(iteration *models.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(iteration)
}
