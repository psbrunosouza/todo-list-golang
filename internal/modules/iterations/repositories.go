package iterations

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

func CreateIteration(iteration *entities.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Create(iteration)
}

func ListIterations(iteration *[]entities.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Find(iteration)
}

func FindIteration(id int, iteration *entities.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.First(iteration, id)
}

func UpdateIteration(id int, iteration *entities.Iteration) (result *gorm.DB) {
	databases.PostgresDB.Model(iteration).Where("id = ?", id).Updates(iteration)
	return FindIteration(id, iteration)
}

func DeleteIteration(iteration *entities.Iteration) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(iteration)
}
