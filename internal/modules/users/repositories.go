package users

import (
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) *gorm.DB
	List(users *[]entities.User) *gorm.DB
	Find(id int, user *entities.User) *gorm.DB
	Update(id int, user *entities.User) *gorm.DB
	Delete(user *entities.User) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user *entities.User) (result *gorm.DB) {
	return r.db.Create(user)
}

func (r *repository) List(users *[]entities.User) (result *gorm.DB) {
	return r.db.Find(users)
}

func (r *repository) Find(id int, user *entities.User) (result *gorm.DB) {
	return r.db.First(user, id)
}

func (r *repository) Update(id int, user *entities.User) (result *gorm.DB) {
	r.db.Model(user).Where("id = ?", id).Updates(user)
	return r.db.First(user, id)
}

func (r *repository) Delete(user *entities.User) (result *gorm.DB) {
	return r.db.Unscoped().Delete(user)
}
