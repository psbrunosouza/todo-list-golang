package subtasks

func CreateSubTaskService(subTask *SubTask) error {
	if result := CreateSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateSubTaskService(subTask *SubTask) error {
	if result := UpdateSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindSubTaskService(id int, subTask *SubTask) error {
	if result := FindSubTask(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListSubTasksService(subTasks *[]SubTask) error {
	if result := ListSubTasks(subTasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteSubTaskService(subTask *SubTask) error {
	if result := DeleteSubTask(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
