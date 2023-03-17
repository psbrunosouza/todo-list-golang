package iterations

import "todo-list/internal/entities"

func CreateIterationService(iteration *entities.Iteration) error {
	if result := CreateIteration(iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteIterationService(iteration *entities.Iteration) error {
	if result := DeleteIteration(iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindIterationService(id int, iteration *entities.Iteration) error {
	if result := FindIteration(id, iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListIterationsService(iteration *[]entities.Iteration) error {
	if result := ListIterations(iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateIterationService(id int, iteration *entities.Iteration) error {
	if result := UpdateIteration(id, iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
