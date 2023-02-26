package subtasks

import (
	"todo-list/internal/databases"

	"gorm.io/gorm"
)

func CreateSubTask(subtask *SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Create(subtask)
}

func ListSubTasks(subtasks *[]SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Find(subtasks)
}

func FindSubTask(id int, subtask *SubTask) (result *gorm.DB) {
	return databases.PostgresDB.First(subtask, id)
}

func UpdateSubTask(subtask *SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Save(subtask)
}

func DeleteSubTask(subtask *SubTask) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(subtask)
}
