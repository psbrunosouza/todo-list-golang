package ratings

import (
	"todo-list/internal/databases"
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

func CreateRating(rating *entities.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Create(rating)
}

func ListRatings(ratings *[]entities.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Find(ratings)
}

func FindRating(id int, rating *entities.Rating) (result *gorm.DB) {
	return databases.PostgresDB.First(rating, id)
}

func UpdateRating(id int, rating *entities.Rating) (result *gorm.DB) {
	databases.PostgresDB.Model(rating).Where("id = ?", id).Updates(rating)
	return FindRating(id, rating)
}

func DeleteRating(rating *entities.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(rating)
}
