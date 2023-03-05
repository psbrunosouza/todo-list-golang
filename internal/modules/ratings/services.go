package ratings

import "todo-list/internal/models"

func CreateRatingService(rating *models.Rating) error {
	if result := CreateRating(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteRatingService(rating *models.Rating) error {
	if result := DeleteRating(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindRatingService(id int, rating *models.Rating) error {
	if result := FindRating(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListRatingsService(ratings *[]models.Rating) error {
	if result := ListRatings(ratings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateRatingService(id int, rating *models.Rating) error {
	if result := UpdateRating(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
