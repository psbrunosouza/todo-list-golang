package tasks

import "todo-list/internal/entities"

type TaskService interface {
	Create(task *entities.Task) error
	List(tasks *[]entities.Task) error
	Find(id int, task *entities.Task) error
	Update(id int, task *entities.Task) error
	Delete(task *entities.Task) error
}

type service struct {
	repo TaskRepository
}

func NewSubTaskService(repo TaskRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(task *entities.Task) error {
	if result := s.repo.Create(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(task *entities.Task) error {
	if result := s.repo.Delete(task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, task *entities.Task) error {
	if result := s.repo.Find(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(tasks *[]entities.Task) error {
	if result := s.repo.List(tasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, task *entities.Task) error {
	if result := s.repo.Update(id, task); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
