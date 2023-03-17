package iterationsettings

import "todo-list/internal/entities"

func CreateIterationSettingService(iterationsetting *entities.IterationSetting) error {
	if result := CreateIterationSetting(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteIterationSettingService(iterationsetting *entities.IterationSetting) error {
	if result := DeleteIterationSetting(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindIterationSettingService(id int, iterationsetting *entities.IterationSetting) error {
	if result := FindIterationSetting(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListIterationSettingsService(iterationsettings *[]entities.IterationSetting) error {
	if result := ListIterationSettings(iterationsettings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateIterationSettingService(id int, iterationsetting *entities.IterationSetting) error {
	if result := UpdateIterationSetting(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
