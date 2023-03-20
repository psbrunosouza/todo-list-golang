package subtasks

import (
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type SubTaskRepository interface {
	Create(subtask *entities.SubTask) *gorm.DB
	List(subtasks *[]entities.SubTask) *gorm.DB
	Find(id int, subtask *entities.SubTask) *gorm.DB
	Update(id int, subtask *entities.SubTask) *gorm.DB
	Delete(subtask *entities.SubTask) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewSubTaskRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(subtask *entities.SubTask) (result *gorm.DB) {
	return r.db.Create(subtask)
}

func (r *repository) List(subtasks *[]entities.SubTask) (result *gorm.DB) {
	return r.db.Find(subtasks)
}

func (r *repository) Find(id int, subtask *entities.SubTask) (result *gorm.DB) {
	return r.db.First(subtask, id)
}

func (r *repository) Update(id int, subtask *entities.SubTask) (result *gorm.DB) {
	r.db.Model(subtask).Where("id = ?", id).Updates(subtask)
	return r.db.First(subtask, id)
}

func (r *repository) Delete(subtask *entities.SubTask) (result *gorm.DB) {
	return r.db.Unscoped().Delete(subtask)
}
