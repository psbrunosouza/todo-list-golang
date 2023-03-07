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
	return databases.PostgresDB.Find(subtasks)
}

func FindSubTask(id int, subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.First(subtask, id)
}

func UpdateSubTask(id int, subtask *models.SubTask) (result *gorm.DB) {
	databases.PostgresDB.Model(subtask).Where("id = ?", id).Updates(subtask)
	return FindSubTask(id, subtask)
}

func DeleteSubTask(subtask *models.SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(subtask)
}

