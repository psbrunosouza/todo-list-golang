package tasks

import (
	taskModel "todo-list/internal/modules/tasks/models"
	taskRepository "todo-list/internal/modules/tasks/repositories"
)

func CreateTaskService(task *taskModel.Task) error {
	if result := taskRepository.CreateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
