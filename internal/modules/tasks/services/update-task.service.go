package task_services

import (
	taskModel "todo-list/internal/modules/tasks/task_models"
	tasks "todo-list/internal/modules/tasks/task_repositories"
)

func UpdateTaskService(task *taskModel.Task) error {
	if result := tasks.UpdateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
