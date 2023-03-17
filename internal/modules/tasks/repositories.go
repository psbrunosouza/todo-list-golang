package tasks

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

func CreateTask(task *entities.Task) (result *gorm.DB) {
	return databases.PostgresDB.Create(task)
}

func ListTasks(tasks *[]entities.Task) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Workspace").Preload("Subtasks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, task_id")
	}).Find(tasks)
}

func FindTask(id int, task *entities.Task) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Workspace").Preload("Subtasks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, task_id")
	}).First(task, id)
}

func UpdateTask(id int, task *entities.Task) (result *gorm.DB) {
	databases.PostgresDB.Model(task).Where("id = ?", id).Updates(task)
	return FindTask(id, task)
}

func DeleteTask(task *entities.Task) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(task)
}
