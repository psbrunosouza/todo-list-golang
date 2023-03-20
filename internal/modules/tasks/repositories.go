package tasks

import (
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *entities.Task) *gorm.DB
	List(tasks *[]entities.Task) *gorm.DB
	Find(id int, task *entities.Task) *gorm.DB
	Update(id int, task *entities.Task) *gorm.DB
	Delete(task *entities.Task) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(task *entities.Task) (result *gorm.DB) {
	return r.db.Create(task)
}

func (r *repository) List(tasks *[]entities.Task) (result *gorm.DB) {
	return r.db.Preload("Workspace").Preload("Subtasks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, task_id")
	}).Find(tasks)
}

func (r *repository) Find(id int, task *entities.Task) (result *gorm.DB) {
	return r.db.Preload("Workspace").Preload("Subtasks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, task_id")
	}).First(task, id)
}

func (r *repository) Update(id int, task *entities.Task) (result *gorm.DB) {
	r.db.Model(task).Where("id = ?", id).Updates(task)
	return r.db.Preload("Workspace").Preload("Subtasks", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, task_id")
	}).First(task, id)
}

func (r *repository) Delete(task *entities.Task) (result *gorm.DB) {
	return r.db.Unscoped().Delete(task)
}
