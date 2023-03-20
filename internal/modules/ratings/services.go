package ratings

import (
	"todo-list/internal/entities"
)

type RatingService interface {
	Create(rating *entities.Rating) error
	List(ratings *[]entities.Rating) error
	Find(id int, rating *entities.Rating) error
	Update(id int, rating *entities.Rating) error
	Delete(rating *entities.Rating) error
}

type service struct {
	repo RatingRepository
}

func NewRatingService(repo RatingRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(rating *entities.Rating) error {
	if result := s.repo.Create(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(rating *entities.Rating) error {
	if result := s.repo.Delete(rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, rating *entities.Rating) error {
	if result := s.repo.Find(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(ratings *[]entities.Rating) error {
	if result := s.repo.List(ratings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, rating *entities.Rating) error {
	if result := s.repo.Update(id, rating); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
