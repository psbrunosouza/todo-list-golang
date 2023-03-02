package tasks

import "todo-list/internal/models"

func CreateTaskService(task *models.Task) error {
	if result := CreateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteTaskService(task *models.Task) error {
	if result := DeleteTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindTaskService(id int, task *models.Task) error {
	if result := FindTask(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListTasksService(tasks *[]models.Task) error {
	if result := ListTasks(tasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateTaskService(task *models.Task) error {
	if result := UpdateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
