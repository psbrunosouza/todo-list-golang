package tasks

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateTask(task *models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Create(task)
}

func ListTasks(tasks *[]models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Subtasks").Preload("Workspace").Preload("Workspace.Tasks").Find(tasks)
}

func FindTask(id int, task *models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Subtasks").Preload("Workspace").Preload("Workspace.Tasks").First(task, id)
}

func UpdateTask(task *models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Save(task)
}

func DeleteTask(task *models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(task)
}
