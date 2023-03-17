package colors

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"
)

var colorRepository ColorRepository = NewColorRepository(databases.PostgresDB)

func CreateColorService(color *entities.Color) error {
	if result := colorRepository.Create(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteColorService(color *entities.Color) error {
	if result := colorRepository.Delete(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindColorService(id int, color *entities.Color) error {
	if result := colorRepository.Find(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListColorsService(colors *[]entities.Color) error {
	if result := colorRepository.List(colors); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateColorService(id int, color *entities.Color) error {
	if result := colorRepository.Update(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
