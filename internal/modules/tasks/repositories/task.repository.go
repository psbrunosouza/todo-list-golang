package task_repositories

import (
	"todo-list/internal/databases"
	taskModels "todo-list/internal/modules/tasks/task_models"

	"gorm.io/gorm"
)

func CreateTask(task *taskModels.Task) (result *gorm.DB) {
	return databases.PostgresDB.Create(task)
}

func ListTasks(tasks *[]taskModels.Task) (result *gorm.DB) {
	return databases.PostgresDB.Find(tasks)
}

func FindTask(id int, task *taskModels.Task) (result *gorm.DB) {
	return databases.PostgresDB.First(task, id)
}

func UpdateTask(task *taskModels.Task) (result *gorm.DB) {
	return databases.PostgresDB.Save(task)
}

func DeleteTask(task *taskModels.Task) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(task)
}
