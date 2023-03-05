package ratings

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateRating(rating *models.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Create(rating)
}

func ListRatings(ratings *[]models.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Find(ratings)
}

func FindRating(id int, rating *models.Rating) (result *gorm.DB) {
	return databases.PostgresDB.First(rating, id)
}

func UpdateRating(id int, rating *models.Rating) (result *gorm.DB) {
	databases.PostgresDB.Model(rating).Where("id = ?", id).Updates(rating)
	return FindRating(id, rating)
}

func DeleteRating(rating *models.Rating) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(rating)
}
