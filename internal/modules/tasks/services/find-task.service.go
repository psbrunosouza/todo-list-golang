package task_services

import (
	task_models "todo-list/internal/modules/tasks/models"
	task_repositories "todo-list/internal/modules/tasks/repositories"
)

func FindTaskService(id int, task *task_models.Task) error {
	if result := task_repositories.FindTask(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
