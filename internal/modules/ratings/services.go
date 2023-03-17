package ratings

import "todo-list/internal/entities"

func CreateRatingService(rating *entities.Rating) error {
	if result := CreateRating(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteRatingService(rating *entities.Rating) error {
	if result := DeleteRating(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindRatingService(id int, rating *entities.Rating) error {
	if result := FindRating(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListRatingsService(ratings *[]entities.Rating) error {
	if result := ListRatings(ratings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateRatingService(id int, rating *entities.Rating) error {
	if result := UpdateRating(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
