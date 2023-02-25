package task_services

import (
	taskModels "todo-list/internal/modules/tasks/task_models"
	taskRepositories "todo-list/internal/modules/tasks/task_repositories"
)

func ListTasksService(tasks *[]taskModels.Task) error {
	if result := taskRepositories.ListTasks(tasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
