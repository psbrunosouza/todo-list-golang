package iterations

import (
	"todo-list/internal/entities"
)

type IterationService interface {
	Create(iteration *entities.Iteration) error
	List(iterations *[]entities.Iteration) error
	Find(id int, iteration *entities.Iteration) error
	Update(id int, iteration *entities.Iteration) error
	Delete(iteration *entities.Iteration) error
}

type service struct {
	repo IterationRepository
}

func NewIterationService(repo IterationRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(iteration *entities.Iteration) error {
	if result := s.repo.Create(iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(iteration *entities.Iteration) error {
	if result := s.repo.Delete(iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, iteration *entities.Iteration) error {
	if result := s.repo.Find(id, iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(iterations *[]entities.Iteration) error {
	if result := s.repo.List(iterations); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, iteration *entities.Iteration) error {
	if result := s.repo.Update(id, iteration); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
