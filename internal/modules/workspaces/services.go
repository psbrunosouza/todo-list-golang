package workspaces

import "todo-list/internal/entities"

type WorkspaceService interface {
	Create(workspace *entities.Workspace) error
	List(workspaces *[]entities.Workspace) error
	Find(id int, workspace *entities.Workspace) error
	Update(id int, workspace *entities.Workspace) error
	Delete(workspace *entities.Workspace) error
}

type service struct {
	repo WorkspaceRepository
}

func NewWorkspaceService(repo WorkspaceRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(workspace *entities.Workspace) error {
	if result := s.repo.Create(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Delete(workspace *entities.Workspace) error {
	if result := s.repo.Delete(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Find(id int, workspace *entities.Workspace) error {
	if result := s.repo.Find(id, workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) List(workspaces *[]entities.Workspace) error {
	if result := s.repo.List(workspaces); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (s *service) Update(id int, workspace *entities.Workspace) error {
	if result := s.repo.Update(id, workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
