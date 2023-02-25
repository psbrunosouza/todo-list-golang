package tasks

import (
	taskModel "todo-list/internal/modules/tasks/models"
	tasks "todo-list/internal/modules/tasks/repositories"
)

func DeleteTaskService(task *taskModel.Task) error {
	if result := tasks.DeleteTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
