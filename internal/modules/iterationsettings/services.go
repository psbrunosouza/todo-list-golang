package iterationsettings

import (
	"todo-list/internal/entities"
)

type IterationSettingService interface {
	Create(iterationsetting *entities.IterationSetting) error
	List(iterationsettings *[]entities.IterationSetting) error
	Find(id int, iterationsetting *entities.IterationSetting) error
	Update(id int, iterationsetting *entities.IterationSetting) error
	Delete(iterationsetting *entities.IterationSetting) error
}

type service struct {
	repo IterationSettingRepository
}

func NewIterationSettingService(repo IterationSettingRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(iterationsetting *entities.IterationSetting) error {
	if result := s.repo.Create(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(iterationsetting *entities.IterationSetting) error {
	if result := s.repo.Delete(iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, iterationsetting *entities.IterationSetting) error {
	if result := s.repo.Find(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(iterationsettings *[]entities.IterationSetting) error {
	if result := s.repo.List(iterationsettings); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, iterationsetting *entities.IterationSetting) error {
	if result := s.repo.Update(id, iterationsetting); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
