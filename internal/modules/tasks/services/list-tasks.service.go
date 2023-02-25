package tasks

import (
	taskModels "todo-list/internal/modules/tasks/models"
	taskRepositories "todo-list/internal/modules/tasks/repositories"
)

func ListTasksService(tasks *[]taskModels.Task) error {
	if result := taskRepositories.ListTasks(tasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
