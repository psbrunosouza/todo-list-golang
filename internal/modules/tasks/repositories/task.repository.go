package task_repositories

import (
	"todo-list/internal/databases"
	task_models "todo-list/internal/modules/tasks/models"

	"gorm.io/gorm"
)

func CreateTask(task *task_models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Create(task)
}

func ListTasks(tasks *[]task_models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Find(tasks)
}

func FindTask(id int, task *task_models.Task) (result *gorm.DB) {
	return databases.PostgresDB.First(task, id)
}

func UpdateTask(task *task_models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Save(task)
}

func DeleteTask(task *task_models.Task) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(task)
}
