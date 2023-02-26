package tasks

import (
	"todo-list/internal/databases"

	"gorm.io/gorm"
)

func CreateTask(task *Task) (result *gorm.DB) {
	return databases.PostgresDB.Create(task)
}

func ListTasks(tasks *[]Task) (result *gorm.DB) {
	return databases.PostgresDB.Find(tasks)
}

func FindTask(id int, task *Task) (result *gorm.DB) {
	return databases.PostgresDB.First(task, id)
}

func UpdateTask(task *Task) (result *gorm.DB) {
	return databases.PostgresDB.Save(task)
}

func DeleteTask(task *Task) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(task)
}
