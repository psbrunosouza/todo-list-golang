package colors

import "todo-list/internal/models"

func CreateColorService(color *models.Color) error {
	if result := CreateColor(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteColorService(color *models.Color) error {
	if result := DeleteColor(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindColorService(id int, color *models.Color) error {
	if result := FindColor(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListColorsService(colors *[]models.Color) error {
	if result := ListColors(colors); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateColorService(id int, color *models.Color) error {
	if result := UpdateColor(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
