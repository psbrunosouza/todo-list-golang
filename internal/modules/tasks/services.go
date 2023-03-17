package tasks

import "todo-list/internal/entities"

func CreateTaskService(task *entities.Task) error {
	if result := CreateTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteTaskService(task *entities.Task) error {
	if result := DeleteTask(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindTaskService(id int, task *entities.Task) error {
	if result := FindTask(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListTasksService(tasks *[]entities.Task) error {
	if result := ListTasks(tasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateTaskService(id int, task *entities.Task) error {
	if result := UpdateTask(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
