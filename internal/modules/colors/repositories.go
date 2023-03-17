package colors

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type ColorRepository interface {
	Create(color *entities.Color) *gorm.DB
	List(colors *[]entities.Color) *gorm.DB
	Find(id int, color *entities.Color) *gorm.DB
	Update(id int, color *entities.Color) *gorm.DB
	Delete(color *entities.Color) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewColorRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(color *entities.Color) (result *gorm.DB) {
	return r.db.Create(color)
}

func (r *repository) List(colors *[]entities.Color) (result *gorm.DB) {
	return r.db.Find(colors)
}

func (r *repository) Find(id int, color *entities.Color) (result *gorm.DB) {
	return r.db.First(color, id)
}

func (r *repository) Update(id int, color *entities.Color) (result *gorm.DB) {
	databases.PostgresDB.Model(color).Where("id = ?", id).Updates(color)
	return r.db.First(color, id)
}

func (r *repository) Delete(color *entities.Color) (result *gorm.DB) {
	return r.db.Unscoped().Delete(color)
}
