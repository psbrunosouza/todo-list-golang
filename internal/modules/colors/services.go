package colors

import (
	"todo-list/internal/entities"
)

type ColorService interface {
	Create(color *entities.Color) error
	List(colors *[]entities.Color) error
	Find(id int, color *entities.Color) error
	Update(id int, color *entities.Color) error
	Delete(color *entities.Color) error
}

type service struct {
	repo ColorRepository
}

func NewColorService(repo ColorRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(color *entities.Color) error {
	if result := s.repo.Create(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(color *entities.Color) error {
	if result := s.repo.Delete(color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, color *entities.Color) error {
	if result := s.repo.Find(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(colors *[]entities.Color) error {
	if result := s.repo.List(colors); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, color *entities.Color) error {
	if result := s.repo.Update(id, color); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
