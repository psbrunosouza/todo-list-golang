package subtasks

import "todo-list/internal/entities"

func CreateSubTaskService(subTask *entities.SubTask) error {
	if result := CreateSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateSubTaskService(id int, subTask *entities.SubTask) error {
	if result := UpdateSubTask(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindSubTaskService(id int, subTask *entities.SubTask) error {
	if result := FindSubTask(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListSubTasksService(subTasks *[]entities.SubTask) error {
	if result := ListSubTasks(subTasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteSubTaskService(subTask *entities.SubTask) error {
	if result := DeleteSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
