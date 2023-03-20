package subtasks

import "todo-list/internal/entities"

type SubTaskService interface {
	Create(subTask *entities.SubTask) error
	List(subTasks *[]entities.SubTask) error
	Find(id int, subTask *entities.SubTask) error
	Update(id int, subTask *entities.SubTask) error
	Delete(subTask *entities.SubTask) error
}

type service struct {
	repo SubTaskRepository
}

func NewSubTaskService(repo SubTaskRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(subTask *entities.SubTask) error {
	if result := s.repo.Create(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, subTask *entities.SubTask) error {
	if result := s.repo.Update(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, subTask *entities.SubTask) error {
	if result := s.repo.Find(id, subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(subTasks *[]entities.SubTask) error {
	if result := s.repo.List(subTasks); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(subTask *entities.SubTask) error {
	if result := s.repo.Delete(subTask); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
