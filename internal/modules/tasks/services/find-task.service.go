package tasks

import (
	taskModel "todo-list/internal/modules/tasks/models"
	tasks "todo-list/internal/modules/tasks/repositories"
)

func FindTaskService(id int, task *taskModel.Task) error {
	if result := tasks.FindTask(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
