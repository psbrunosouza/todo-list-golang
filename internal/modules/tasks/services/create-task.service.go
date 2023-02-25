package task_services

import (
	taskModel "todo-list/internal/modules/tasks/task_models"
	taskRepository "todo-list/internal/modules/tasks/task_repositories"
)

func CreateTaskService(task *taskModel.Task) error {
	if result := taskRepository.CreateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
