package subtasks

import "todo-list/internal/models"

func CreateSubTaskService(subTask *models.SubTask) error {
	if result := CreateSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateSubTaskService(id int, subTask *models.SubTask) error {
	if result := UpdateSubTask(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindSubTaskService(id int, subTask *models.SubTask) error {
	if result := FindSubTask(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListSubTasksService(subTasks *[]models.SubTask) error {
	if result := ListSubTasks(subTasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteSubTaskService(subTask *models.SubTask) error {
	if result := DeleteSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func MarkAsDoneService(id int, subtask *models.SubTask) error {
	if result := MarkAsDone(id, subtask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
