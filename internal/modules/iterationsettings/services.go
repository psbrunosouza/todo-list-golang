package iterationsettings

import "todo-list/internal/models"

func CreateIterationSettingService(iterationsetting *models.IterationSetting) error {
	if result := CreateIterationSetting(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteIterationSettingService(iterationsetting *models.IterationSetting) error {
	if result := DeleteIterationSetting(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindIterationSettingService(id int, iterationsetting *models.IterationSetting) error {
	if result := FindIterationSetting(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListIterationSettingsService(iterationsettings *[]models.IterationSetting) error {
	if result := ListIterationSettings(iterationsettings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateIterationSettingService(id int, iterationsetting *models.IterationSetting) error {
	if result := UpdateIterationSetting(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
