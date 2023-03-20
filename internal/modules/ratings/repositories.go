package ratings

import (
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(rating *entities.Rating) *gorm.DB
	List(ratings *[]entities.Rating) *gorm.DB
	Find(id int, rating *entities.Rating) *gorm.DB
	Update(id int, rating *entities.Rating) *gorm.DB
	Delete(rating *entities.Rating) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(rating *entities.Rating) (result *gorm.DB) {
	return r.db.Create(rating)
}

func (r *repository) List(ratings *[]entities.Rating) (result *gorm.DB) {
	return r.db.Find(ratings)
}

func (r *repository) Find(id int, rating *entities.Rating) (result *gorm.DB) {
	return r.db.First(rating, id)
}

func (r *repository) Update(id int, rating *entities.Rating) (result *gorm.DB) {
	r.db.Model(rating).Where("id = ?", id).Updates(rating)
	return r.db.First(rating, id)
}

func (r *repository) Delete(rating *entities.Rating) (result *gorm.DB) {
	return r.db.Unscoped().Delete(rating)
}
