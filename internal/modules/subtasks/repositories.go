package subtasks

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateSubTask(subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Create(subtask)
}

func ListSubTasks(subtasks *[]models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Task").Find(subtasks)
}

func FindSubTask(id int, subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Task").First(subtask, id)
}

func UpdateSubTask(subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Save(subtask)
}

func DeleteSubTask(subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(subtask)
}
