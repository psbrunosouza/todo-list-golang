package tasks

import (
	"todo-list/internal/databases"
	taskModels "todo-list/internal/modules/tasks/models"
)

func CreateTask(task *taskModels.Task) error {
	db := databases.PostgresDB.Create(task)
	return db.Error
}
