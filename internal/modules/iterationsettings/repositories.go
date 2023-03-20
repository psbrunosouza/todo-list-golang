package iterationsettings

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type IterationSettingRepository interface {
	Create(iterationSetting *entities.IterationSetting) *gorm.DB
	List(iterationSettings *[]entities.IterationSetting) *gorm.DB
	Find(id int, iterationSetting *entities.IterationSetting) *gorm.DB
	Update(id int, iterationSetting *entities.IterationSetting) *gorm.DB
	Delete(iterationSetting *entities.IterationSetting) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewIterationSettingRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return r.db.Create(iterationSetting)
}

func (r *repository) List(iterationSetting *[]entities.IterationSetting) (result *gorm.DB) {
	return r.db.Find(iterationSetting)
}

func (r *repository) Find(id int, iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return r.db.First(iterationSetting, id)
}

func (r *repository) Update(id int, iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	databases.PostgresDB.Model(iterationSetting).Where("id = ?", id).Updates(iterationSetting)
	return r.db.First(iterationSetting, id)
}

func (r *repository) Delete(iterationSetting *entities.IterationSetting) (result *gorm.DB) {
	return r.db.Unscoped().Delete(iterationSetting)
}
