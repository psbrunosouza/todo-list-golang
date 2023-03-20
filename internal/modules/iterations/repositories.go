package iterations

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type IterationRepository interface {
	Create(iteration *entities.Iteration) *gorm.DB
	List(iterations *[]entities.Iteration) *gorm.DB
	Find(id int, iteration *entities.Iteration) *gorm.DB
	Update(id int, iteration *entities.Iteration) *gorm.DB
	Delete(iteration *entities.Iteration) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewIterationRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(iteration *entities.Iteration) (result *gorm.DB) {
	return r.db.Create(iteration)
}

func (r *repository) List(iteration *[]entities.Iteration) (result *gorm.DB) {
	return r.db.Find(iteration)
}

func (r *repository) Find(id int, iteration *entities.Iteration) (result *gorm.DB) {
	return r.db.First(iteration, id)
}

func (r *repository) Update(id int, iteration *entities.Iteration) (result *gorm.DB) {
	databases.PostgresDB.Model(iteration).Where("id = ?", id).Updates(iteration)
	return r.db.First(iteration, id)
}

func (r *repository) Delete(iteration *entities.Iteration) (result *gorm.DB) {
	return r.db.Unscoped().Delete(iteration)
}
