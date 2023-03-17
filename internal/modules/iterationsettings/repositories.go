package iterationsettings

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

func CreateIterationSetting(iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return databases.PostgresDB.Create(iterationSetting)
}

func ListIterationSettings(iterationSetting *[]entities.IterationSetting) (result *gorm.DB) {
	return databases.PostgresDB.Find(iterationSetting)
}

func FindIterationSetting(id int, iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return databases.PostgresDB.First(iterationSetting, id)
}

func UpdateIterationSetting(id int, iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	databases.PostgresDB.Model(iterationSetting).Where("id = ?", id).Updates(iterationSetting)
	return FindIterationSetting(id, iterationSetting)
}

func DeleteIterationSetting(iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(iterationSetting)
}
