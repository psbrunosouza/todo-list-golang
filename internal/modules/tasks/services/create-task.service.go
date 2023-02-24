package tasks

import (
	"log"
	taskModel "todo-list/internal/modules/tasks/models"
	taskRepository "todo-list/internal/modules/tasks/repositories"
)

func CreateTaskService(task *taskModel.Task) {
	if err := taskRepository.CreateTask(task); err != nil {
		log.Fatalln("Error during task creation")
	}
}
