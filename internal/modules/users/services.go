package users

import "todo-list/internal/entities"

type UserService interface {
	Create(user *entities.User) error
	List(users *[]entities.User) error
	Find(id int, user *entities.User) error
	Update(id int, user *entities.User) error
	Delete(user *entities.User) error
}

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(user *entities.User) error {
	if result := s.repo.Create(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(user *entities.User) error {
	if result := s.repo.Delete(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, user *entities.User) error {
	if result := s.repo.Find(id, user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(users *[]entities.User) error {
	if result := s.repo.List(users); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, user *entities.User) error {
	if result := s.repo.Update(id, user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
